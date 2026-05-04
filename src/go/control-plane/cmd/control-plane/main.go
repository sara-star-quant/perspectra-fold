package main

import (
	"log"
	"os"
	"time"

	"github.com/peterz/multidimensional-transformation/control-plane/internal/apiserver"
	"github.com/peterz/multidimensional-transformation/control-plane/internal/coreclient"
	"github.com/peterz/multidimensional-transformation/control-plane/internal/store"
)

const version = "0.1.0"

func main() {
	addr := getenv("MDQC_CONTROL_ADDR", ":8080")
	dataPath := getenv("MDQC_CONTROL_DATA_PATH", "data/control-plane.json")

	coreAddr := getenv("MDQC_CORE_ADDR", "127.0.0.1:50051")
	coreTimeout := getenvDuration("MDQC_CORE_TIMEOUT", 250*time.Millisecond)

	telemetryStore, err := store.NewFileStore(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	core, err := coreclient.New(coreAddr, coreTimeout)
	if err != nil {
		log.Printf("core compute unavailable (%s), falling back to local compute", err)
		core = coreclient.NewLocal()
	}
	defer func() {
		if err := core.Close(); err != nil {
			log.Printf("core client close error: %s", err)
		}
	}()

	server := apiserver.New(addr, version, telemetryStore, core)

	log.Printf("control-plane listening on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getenvDuration(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		parsed, err := time.ParseDuration(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}
