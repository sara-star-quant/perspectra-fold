# Source Layout

- `python/`: research and analysis utilities (Python 3.12+).
- `rust/mdqc-core/`: core compute and cryptographic primitives (Rust).
- `rust/compute-service/`: gRPC compute service (Rust).
- `go/control-plane/`: control plane service (Go).
- `../proto/`: shared protobuf definitions.

## Quick Start (Local)

Python (research utilities):

```
python -m venv .venv
. .venv/bin/activate
pip install -e src/python
```

Rust (core crate):

```
cargo test --manifest-path src/rust/mdqc-core/Cargo.toml
```

Rust (compute service):

```
MDQC_CORE_ADDR=127.0.0.1:50051 cargo run --manifest-path src/rust/compute-service/Cargo.toml
```

Go (control plane):

```
cd src/go/control-plane
MDQC_CONTROL_DATA_PATH=../../data/control-plane.json \
MDQC_CORE_ADDR=127.0.0.1:50051 \
go run ./cmd/control-plane
```

## Control Plane Notes

- Endpoints: `GET/POST /api/v1/links`, `GET/POST /api/v1/metrics`, `GET /api/v1/status`, `GET /api/v1/routes`.
- Telemetry persists to JSON at `MDQC_CONTROL_DATA_PATH`.
- Core compute gRPC target is configured via `MDQC_CORE_ADDR` and `MDQC_CORE_TIMEOUT`.
- Routing policies configured via `MDQC_ROUTE_POLICY` and `MDQC_ROUTE_*` overrides.

Example: posting link telemetry with loss/latency/jitter fields:

```bash
curl -X POST http://127.0.0.1:8080/api/v1/links \
  -H 'Content-Type: application/json' \
  -d '[{"id":"qkd-1","state":"up","qber":0.032,"skr_kbps":12.5,"attenuation_db":3.2,"latency_ms":8.4,"jitter_ms":0.6}]'
```

To regenerate Go gRPC stubs from `proto/core_compute.proto`:

```
protoc -I proto \
  --go_out=src/go/control-plane \
  --go_opt=module=github.com/peterz/multidimensional-transformation/control-plane \
  --go-grpc_out=src/go/control-plane \
  --go-grpc_opt=module=github.com/peterz/multidimensional-transformation/control-plane \
  proto/core_compute.proto
```

To run the compute service + control plane and fetch metrics in one step:

```
make verify-metrics
```
