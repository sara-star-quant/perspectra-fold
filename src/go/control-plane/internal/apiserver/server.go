package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/peterz/multidimensional-transformation/control-plane/internal/coreclient"
	"github.com/peterz/multidimensional-transformation/control-plane/internal/models"
	"github.com/peterz/multidimensional-transformation/control-plane/internal/routing"
	"github.com/peterz/multidimensional-transformation/control-plane/internal/store"
)

const maxPayloadBytes = 1 << 20

type Server struct {
	addr      string
	mux       *http.ServeMux
	startTime time.Time
	version   string
	core      coreclient.Client
	store     *store.Store
}

func New(addr string, version string, st *store.Store, core coreclient.Client) *Server {
	mux := http.NewServeMux()
	srv := &Server{
		addr:      addr,
		mux:       mux,
		startTime: time.Now().UTC(),
		version:   version,
		core:      core,
		store:     st,
	}

	mux.HandleFunc("/api/v1/status", srv.handleStatus)
	mux.HandleFunc("/api/v1/links", srv.handleLinks)
	mux.HandleFunc("/api/v1/metrics", srv.handleMetrics)
	mux.HandleFunc("/api/v1/routes", srv.handleRoutes)

	return srv
}

func (s *Server) ListenAndServe() error {
	httpServer := &http.Server{
		Addr:              s.addr,
		Handler:           s.mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	return httpServer.ListenAndServe()
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	status := models.Status{
		Service:       "control-plane",
		Version:       s.version,
		UptimeSeconds: int64(time.Since(s.startTime).Seconds()),
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
	}
	writeJSON(w, http.StatusOK, status)
}

func (s *Server) handleLinks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, s.store.Links())
	case http.MethodPost:
		var payload []models.Link
		if err := readJSON(w, r, &payload); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		if err := s.store.UpsertLinks(payload); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		writeJSON(w, http.StatusOK, s.store.Links())
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		metrics := s.store.Metrics()
		metrics = s.withQualityScore(r.Context(), metrics)
		metrics = s.withRouteScore(metrics, s.store.Links())
		writeJSON(w, http.StatusOK, metrics)
	case http.MethodPost:
		var payload []models.Metric
		if err := readJSON(w, r, &payload); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		if err := s.store.UpsertMetrics(payload); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		metrics := s.store.Metrics()
		metrics = s.withQualityScore(r.Context(), metrics)
		metrics = s.withRouteScore(metrics, s.store.Links())
		writeJSON(w, http.StatusOK, metrics)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) withQualityScore(ctx context.Context, metrics []models.Metric) []models.Metric {
	if s.core == nil {
		return metrics
	}

	var qber float64
	var skr float64
	var hasQber bool
	var hasSkr bool
	var qualityIndex int
	var hasQuality bool

	for i, metric := range metrics {
		switch metric.Name {
		case "qber":
			qber = metric.Value
			hasQber = true
		case "skr":
			skr = metric.Value
			hasSkr = true
		case "quality_score":
			qualityIndex = i
			hasQuality = true
		}
	}

	if !hasQber || !hasSkr {
		return metrics
	}

	current := []float64{1 - qber, skr}
	baseline := []float64{1, 15}
	score, err := s.core.CosineSimilarity(ctx, current, baseline)
	if err != nil {
		return metrics
	}

	if hasQuality {
		metrics[qualityIndex].Value = score
		metrics[qualityIndex].Unit = "ratio"
		return metrics
	}

	return append(metrics, models.Metric{Name: "quality_score", Value: score, Unit: "ratio"})
}

func (s *Server) withRouteScore(metrics []models.Metric, links []models.Link) []models.Metric {
	cfg := routing.ConfigFromEnv()
	score, ok := routing.BestRouteScoreWithConfig(cfg, links)
	if !ok {
		return metrics
	}

	for i, metric := range metrics {
		if metric.Name == "route_score" {
			metrics[i].Value = score
			metrics[i].Unit = "ratio"
			return metrics
		}
	}

	return append(metrics, models.Metric{Name: "route_score", Value: score, Unit: "ratio"})
}

func readJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxPayloadBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return err
	}
	if dec.More() {
		return errors.New("multiple JSON values provided")
	}
	return nil
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, errorResponse{Error: err.Error()})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
