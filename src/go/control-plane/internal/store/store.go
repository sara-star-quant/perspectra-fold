package store

import (
    "encoding/json"
    "errors"
    "os"
    "path/filepath"
    "sync"
    "time"

    "github.com/peterz/multidimensional-transformation/control-plane/internal/models"
)

var ErrEmptyPayload = errors.New("empty payload")

const defaultState = "up"

// Store keeps control-plane telemetry in memory and persists it to disk.
type Store struct {
    mu      sync.RWMutex
    path    string
    links   []models.Link
    metrics []models.Metric
}

type snapshot struct {
    Links     []models.Link   `json:"links"`
    Metrics   []models.Metric `json:"metrics"`
    UpdatedAt string          `json:"updated_at"`
}

func NewFileStore(path string) (*Store, error) {
    st := &Store{path: path}
    if err := st.load(); err != nil {
        return nil, err
    }
    return st, nil
}

func (s *Store) Links() []models.Link {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return append([]models.Link(nil), s.links...)
}

func (s *Store) Metrics() []models.Metric {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return append([]models.Metric(nil), s.metrics...)
}

func (s *Store) UpsertLinks(links []models.Link) error {
    if len(links) == 0 {
        return ErrEmptyPayload
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    index := make(map[string]int, len(s.links))
    for i, link := range s.links {
        index[link.ID] = i
    }

    for _, link := range links {
        if link.ID == "" {
            return errors.New("link id is required")
        }
        if link.State == "" {
            link.State = defaultState
        }

        if idx, ok := index[link.ID]; ok {
            s.links[idx] = link
        } else {
            s.links = append(s.links, link)
            index[link.ID] = len(s.links) - 1
        }
    }

    return s.persistLocked()
}

func (s *Store) UpsertMetrics(metrics []models.Metric) error {
    if len(metrics) == 0 {
        return ErrEmptyPayload
    }

    s.mu.Lock()
    defer s.mu.Unlock()

    index := make(map[string]int, len(s.metrics))
    for i, metric := range s.metrics {
        index[metric.Name] = i
    }

    for _, metric := range metrics {
        if metric.Name == "" {
            return errors.New("metric name is required")
        }

        if idx, ok := index[metric.Name]; ok {
            s.metrics[idx] = metric
        } else {
            s.metrics = append(s.metrics, metric)
            index[metric.Name] = len(s.metrics) - 1
        }
    }

    return s.persistLocked()
}

func (s *Store) load() error {
    if s.path == "" {
        return errors.New("store path is required")
    }

    file, err := os.Open(s.path)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            s.links = defaultLinks()
            s.metrics = defaultMetrics()
            return s.persistLocked()
        }
        return err
    }
    defer file.Close()

    var snap snapshot
    if err := json.NewDecoder(file).Decode(&snap); err != nil {
        return err
    }

    s.links = snap.Links
    s.metrics = snap.Metrics
    if len(s.links) == 0 {
        s.links = defaultLinks()
    }
    if len(s.metrics) == 0 {
        s.metrics = defaultMetrics()
    }

    return nil
}

func (s *Store) persistLocked() error {
    snap := snapshot{
        Links:     s.links,
        Metrics:   s.metrics,
        UpdatedAt: time.Now().UTC().Format(time.RFC3339),
    }

    data, err := json.MarshalIndent(snap, "", "  ")
    if err != nil {
        return err
    }

    dir := filepath.Dir(s.path)
    if err := os.MkdirAll(dir, 0o755); err != nil {
        return err
    }

    tmp, err := os.CreateTemp(dir, "mdqc-control-*.json")
    if err != nil {
        return err
    }
    if _, err := tmp.Write(data); err != nil {
        tmp.Close()
        return err
    }
    if err := tmp.Close(); err != nil {
        return err
    }

    return os.Rename(tmp.Name(), s.path)
}

func defaultLinks() []models.Link {
    return []models.Link{
        {ID: "qkd-1", State: "up", QBER: 0.032, SKR: 12.5},
        {ID: "qkd-2", State: "degraded", QBER: 0.071, SKR: 6.8},
    }
}

func defaultMetrics() []models.Metric {
    return []models.Metric{
        {Name: "qber", Value: 0.045, Unit: "ratio"},
        {Name: "skr", Value: 10.2, Unit: "kbps"},
    }
}
