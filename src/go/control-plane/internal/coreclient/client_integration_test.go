//go:build integration
// +build integration

package coreclient

import (
	"bytes"
	"context"
	"errors"
	"math"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestCosineSimilarityGRPC(t *testing.T) {
	if _, err := exec.LookPath("cargo"); err != nil {
		t.Skip("cargo not available")
	}
	if _, err := exec.LookPath("protoc"); err != nil {
		t.Skip("protoc not available")
	}

	root, err := repoRoot()
	if err != nil {
		t.Fatal(err)
	}

	addr := "127.0.0.1:50052"
	serviceCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	output := &bytes.Buffer{}
	cmd := exec.CommandContext(
		serviceCtx,
		"cargo",
		"run",
		"--quiet",
		"--manifest-path",
		filepath.Join(root, "src", "rust", "compute-service", "Cargo.toml"),
	)
	cmd.Env = append(os.Environ(), "MDQC_CORE_ADDR="+addr)
	cmd.Stdout = output
	cmd.Stderr = output

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start compute service: %v", err)
	}

	t.Cleanup(func() {
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
		if output.Len() > 0 {
			t.Logf("compute service output:\n%s", output.String())
		}
	})

	if err := waitForPort(serviceCtx, addr); err != nil {
		t.Fatalf("compute service not ready: %v", err)
	}

	client, err := New(addr, 2*time.Second)
	if err != nil {
		t.Fatalf("failed to connect to compute service: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()

	score, err := client.CosineSimilarity(context.Background(), []float64{1, 2, 3}, []float64{1, 2, 3})
	if err != nil {
		t.Fatalf("cosine similarity call failed: %v", err)
	}
	if math.Abs(score-1.0) > 1e-9 {
		t.Fatalf("expected score ~1.0, got %.12f", score)
	}
}

func repoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for i := 0; i < 8; i++ {
		if fileExists(filepath.Join(dir, "LICENSE")) && fileExists(filepath.Join(dir, "proto", "core_compute.proto")) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", errors.New("repo root not found")
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func waitForPort(ctx context.Context, addr string) error {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		conn, err := net.DialTimeout("tcp", addr, 200*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
