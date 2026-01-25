package apiserver

import (
	"net/http"

	"github.com/peterz/multidimensional-transformation/control-plane/internal/routing"
)

type routeScore struct {
	ID    string  `json:"id"`
	Score float64 `json:"score"`
}

type routesResponse struct {
	Policy string       `json:"policy"`
	Scores []routeScore `json:"scores"`
}

func (s *Server) handleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cfg := routing.ConfigFromEnv()
	links := s.store.Links()
	scores := make([]routeScore, 0, len(links))
	for _, link := range links {
		score, ok := routing.ScoreLinkWithConfig(cfg, link)
		if !ok {
			continue
		}
		scores = append(scores, routeScore{ID: link.ID, Score: score})
	}

	writeJSON(w, http.StatusOK, routesResponse{
		Policy: routing.PolicyName(cfg),
		Scores: scores,
	})
}
