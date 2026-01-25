package routing

import (
	"testing"

	"github.com/peterz/multidimensional-transformation/control-plane/internal/models"
)

func TestScoreLinkBasic(t *testing.T) {
	link := models.Link{
		ID:    "qkd-1",
		State: "up",
		QBER:  0.032,
		SKR:   12.5,
	}

	score, ok := ScoreLink(link)
	if !ok {
		t.Fatal("expected score to be valid")
	}
	if score <= 0 || score > 1 {
		t.Fatalf("unexpected score %.4f", score)
	}
}

func TestScoreLinkRejectsHighQBER(t *testing.T) {
	link := models.Link{
		ID:    "qkd-2",
		State: "up",
		QBER:  0.30,
		SKR:   10,
	}

	if _, ok := ScoreLink(link); ok {
		t.Fatal("expected high QBER link to be rejected")
	}
}

func TestBestRouteScoreSelectsHighest(t *testing.T) {
	links := []models.Link{
		{ID: "down", State: "down", QBER: 0.02, SKR: 20},
		{ID: "good", State: "up", QBER: 0.03, SKR: 15},
		{ID: "degraded", State: "degraded", QBER: 0.04, SKR: 12},
	}

	score, ok := BestRouteScore(links)
	if !ok {
		t.Fatal("expected a valid route score")
	}
	if score <= 0 {
		t.Fatalf("unexpected score %.4f", score)
	}
}

func TestScoreLinkWithCustomConfig(t *testing.T) {
	cfg := DefaultConfig()
	cfg.QBERMax = 0.05
	cfg.TargetSKR = 20
	cfg.PolicyName = "security-first"

	link := models.Link{
		ID:    "qkd-3",
		State: "up",
		QBER:  0.04,
		SKR:   10,
	}

	score, ok := ScoreLinkWithConfig(cfg, link)
	if !ok {
		t.Fatal("expected score to be valid with custom config")
	}
	if score <= 0 {
		t.Fatalf("unexpected score %.4f", score)
	}
}

func TestPolicyNameFallback(t *testing.T) {
	cfg := DefaultConfig()
	cfg.PolicyName = ""

	if got := PolicyName(cfg); got != "custom" {
		t.Fatalf("expected fallback policy name, got %q", got)
	}
}

func TestConfigFromEnvAppliesPolicy(t *testing.T) {
	t.Setenv("MDQC_ROUTE_POLICY", "security-first")

	cfg := ConfigFromEnv()
	if cfg.WeightQBER < 0.59 {
		t.Fatalf("expected security-first weights to apply, got %.2f", cfg.WeightQBER)
	}
	if cfg.QBERMax > 0.09 {
		t.Fatalf("expected stricter QBER max, got %.2f", cfg.QBERMax)
	}
}
