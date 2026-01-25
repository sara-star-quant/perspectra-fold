#!/usr/bin/env bash
set -euo pipefail

export PATH="$HOME/.cargo/bin:$PATH"

CORE_ADDR="${MDQC_CORE_ADDR:-127.0.0.1:50051}"
CONTROL_ADDR="${MDQC_CONTROL_ADDR:-127.0.0.1:8080}"
DATA_PATH="${MDQC_CONTROL_DATA_PATH:-/tmp/mdqc-control-plane.json}"

core_host="${CORE_ADDR%:*}"
core_port="${CORE_ADDR##*:}"

cleanup() {
  if [[ -n "${CONTROL_PID:-}" ]]; then
    kill "${CONTROL_PID}" 2>/dev/null || true
  fi
  if [[ -n "${CORE_PID:-}" ]]; then
    kill "${CORE_PID}" 2>/dev/null || true
  fi
  wait 2>/dev/null || true
}
trap cleanup EXIT

MDQC_CORE_ADDR="${CORE_ADDR}" cargo run --manifest-path src/rust/compute-service/Cargo.toml >/tmp/mdqc-core.log 2>&1 &
CORE_PID=$!

core_ready=false
for _ in $(seq 1 60); do
  if nc -z "${core_host}" "${core_port}" >/dev/null 2>&1; then
    core_ready=true
    break
  fi
  sleep 1
done

if [[ "${core_ready}" != "true" ]]; then
  echo "compute service did not start on ${CORE_ADDR}" >&2
  exit 1
fi

(
  cd src/go/control-plane
  MDQC_CONTROL_ADDR="${CONTROL_ADDR}" \
    MDQC_CORE_ADDR="${CORE_ADDR}" \
    MDQC_CONTROL_DATA_PATH="${DATA_PATH}" \
    go run ./cmd/control-plane
) >/tmp/mdqc-control.log 2>&1 &
CONTROL_PID=$!

control_ready=false
for _ in $(seq 1 30); do
  if curl -sf "http://${CONTROL_ADDR}/api/v1/status" >/dev/null; then
    control_ready=true
    break
  fi
  sleep 1
done

if [[ "${control_ready}" != "true" ]]; then
  echo "control plane did not start on ${CONTROL_ADDR}" >&2
  exit 1
fi

curl -s "http://${CONTROL_ADDR}/api/v1/metrics"
echo
