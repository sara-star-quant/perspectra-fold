package models

type Status struct {
	Service       string `json:"service"`
	Version       string `json:"version"`
	UptimeSeconds int64  `json:"uptime_seconds"`
	Timestamp     string `json:"timestamp"`
}

type Link struct {
	ID            string  `json:"id"`
	State         string  `json:"state"`
	QBER          float64 `json:"qber"`
	SKR           float64 `json:"skr_kbps"`
	AttenuationDB float64 `json:"attenuation_db,omitempty"`
	LatencyMs     float64 `json:"latency_ms,omitempty"`
	JitterMs      float64 `json:"jitter_ms,omitempty"`
}

type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}
