package routing

import (
	"os"
	"strconv"
	"strings"

	"github.com/peterz/multidimensional-transformation/control-plane/internal/models"
)

const (
	defaultQBERMax          = 0.24
	defaultMinSKR           = 1.0
	defaultTargetSKR        = 15.0
	defaultMaxAttenuationDB = 12.0
	defaultMaxLatencyMs     = 50.0
	defaultMaxJitterMs      = 5.0

	defaultWeightQBER        = 0.5
	defaultWeightSKR         = 0.3
	defaultWeightAttenuation = 0.1
	defaultWeightLatency     = 0.05
	defaultWeightJitter      = 0.05

	defaultDegradedPenalty = 0.85
)

type ScoringConfig struct {
	QBERMax          float64
	MinSKR           float64
	TargetSKR        float64
	MaxAttenuationDB float64
	MaxLatencyMs     float64
	MaxJitterMs      float64

	WeightQBER        float64
	WeightSKR         float64
	WeightAttenuation float64
	WeightLatency     float64
	WeightJitter      float64

	DegradedPenalty float64
	PolicyName      string
}

func DefaultConfig() ScoringConfig {
	return ScoringConfig{
		QBERMax:           defaultQBERMax,
		MinSKR:            defaultMinSKR,
		TargetSKR:         defaultTargetSKR,
		MaxAttenuationDB:  defaultMaxAttenuationDB,
		MaxLatencyMs:      defaultMaxLatencyMs,
		MaxJitterMs:       defaultMaxJitterMs,
		WeightQBER:        defaultWeightQBER,
		WeightSKR:         defaultWeightSKR,
		WeightAttenuation: defaultWeightAttenuation,
		WeightLatency:     defaultWeightLatency,
		WeightJitter:      defaultWeightJitter,
		DegradedPenalty:   defaultDegradedPenalty,
		PolicyName:        "balanced",
	}
}

func ConfigFromEnv() ScoringConfig {
	cfg := DefaultConfig()

	cfg.PolicyName = envString("MDQC_ROUTE_POLICY", cfg.PolicyName)

	switch strings.ToLower(strings.TrimSpace(cfg.PolicyName)) {
	case "security-first":
		cfg.WeightQBER = 0.6
		cfg.WeightSKR = 0.3
		cfg.WeightAttenuation = 0.05
		cfg.WeightLatency = 0.03
		cfg.WeightJitter = 0.02
		cfg.QBERMax = 0.08
		cfg.MinSKR = 5
	case "latency-first":
		cfg.WeightQBER = 0.3
		cfg.WeightSKR = 0.2
		cfg.WeightAttenuation = 0.05
		cfg.WeightLatency = 0.35
		cfg.WeightJitter = 0.1
		cfg.QBERMax = 0.2
		cfg.MinSKR = 2
	case "availability-first":
		cfg.WeightQBER = 0.4
		cfg.WeightSKR = 0.4
		cfg.WeightAttenuation = 0.05
		cfg.WeightLatency = 0.1
		cfg.WeightJitter = 0.05
		cfg.QBERMax = 0.24
		cfg.MinSKR = 1
	}

	cfg.QBERMax = envFloat("MDQC_ROUTE_QBER_MAX", cfg.QBERMax)
	cfg.MinSKR = envFloat("MDQC_ROUTE_MIN_SKR", cfg.MinSKR)
	cfg.TargetSKR = envFloat("MDQC_ROUTE_TARGET_SKR", cfg.TargetSKR)
	cfg.MaxAttenuationDB = envFloat("MDQC_ROUTE_MAX_ATTENUATION_DB", cfg.MaxAttenuationDB)
	cfg.MaxLatencyMs = envFloat("MDQC_ROUTE_MAX_LATENCY_MS", cfg.MaxLatencyMs)
	cfg.MaxJitterMs = envFloat("MDQC_ROUTE_MAX_JITTER_MS", cfg.MaxJitterMs)
	cfg.WeightQBER = envFloat("MDQC_ROUTE_WEIGHT_QBER", cfg.WeightQBER)
	cfg.WeightSKR = envFloat("MDQC_ROUTE_WEIGHT_SKR", cfg.WeightSKR)
	cfg.WeightAttenuation = envFloat("MDQC_ROUTE_WEIGHT_ATTENUATION", cfg.WeightAttenuation)
	cfg.WeightLatency = envFloat("MDQC_ROUTE_WEIGHT_LATENCY", cfg.WeightLatency)
	cfg.WeightJitter = envFloat("MDQC_ROUTE_WEIGHT_JITTER", cfg.WeightJitter)
	cfg.DegradedPenalty = envFloat("MDQC_ROUTE_DEGRADED_PENALTY", cfg.DegradedPenalty)

	return cfg
}

func PolicyName(cfg ScoringConfig) string {
	if strings.TrimSpace(cfg.PolicyName) == "" {
		return "custom"
	}
	return strings.TrimSpace(cfg.PolicyName)
}

// ScoreLink returns a normalized score in [0,1] for a single link.
// Lower QBER, higher SKR, and lower loss/latency/jitter score better.
func ScoreLink(link models.Link) (float64, bool) {
	return ScoreLinkWithConfig(DefaultConfig(), link)
}

func ScoreLinkWithConfig(cfg ScoringConfig, link models.Link) (float64, bool) {
	state := strings.ToLower(strings.TrimSpace(link.State))
	if state == "down" {
		return 0, false
	}
	if link.QBER <= 0 || link.SKR <= 0 {
		return 0, false
	}
	if link.QBER > cfg.QBERMax || link.SKR < cfg.MinSKR {
		return 0, false
	}

	score := 0.0
	weightSum := 0.0

	qberScore := clamp01(1 - link.QBER/cfg.QBERMax)
	score += cfg.WeightQBER * qberScore
	weightSum += cfg.WeightQBER

	skrScore := clamp01(link.SKR / cfg.TargetSKR)
	score += cfg.WeightSKR * skrScore
	weightSum += cfg.WeightSKR

	if link.AttenuationDB > 0 {
		lossScore := clamp01(1 - link.AttenuationDB/cfg.MaxAttenuationDB)
		score += cfg.WeightAttenuation * lossScore
		weightSum += cfg.WeightAttenuation
	}

	if link.LatencyMs > 0 {
		latScore := clamp01(1 - link.LatencyMs/cfg.MaxLatencyMs)
		score += cfg.WeightLatency * latScore
		weightSum += cfg.WeightLatency
	}

	if link.JitterMs > 0 {
		jitterScore := clamp01(1 - link.JitterMs/cfg.MaxJitterMs)
		score += cfg.WeightJitter * jitterScore
		weightSum += cfg.WeightJitter
	}

	if weightSum == 0 {
		return 0, false
	}

	score = score / weightSum
	if state == "degraded" {
		score *= cfg.DegradedPenalty
	}

	return clamp01(score), true
}

// BestRouteScore returns the best available score across candidate links.
func BestRouteScore(links []models.Link) (float64, bool) {
	return BestRouteScoreWithConfig(DefaultConfig(), links)
}

func BestRouteScoreWithConfig(cfg ScoringConfig, links []models.Link) (float64, bool) {
	best := 0.0
	ok := false

	for _, link := range links {
		score, valid := ScoreLinkWithConfig(cfg, link)
		if !valid {
			continue
		}
		if !ok || score > best {
			best = score
			ok = true
		}
	}

	return best, ok
}

func clamp01(value float64) float64 {
	if value < 0 {
		return 0
	}
	if value > 1 {
		return 1
	}
	return value
}

func envFloat(key string, fallback float64) float64 {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}
	return parsed
}

func envString(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
