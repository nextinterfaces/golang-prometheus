# Golang Prometheus Monitoring Demo

This project demonstrates how to instrument a basic Go HTTP server with Prometheus metrics and visualize them using Grafana.

---

## Project Structure

- `main.go`: Basic Gin-based HTTP server exposing:
  - `/ping` – simple health check
  - `/metrics` – Prometheus endpoint
- `prometheus.yml`: Configuration file for Prometheus to scrape the Go app
- `docker-compose.yml`: Spins up Prometheus and Grafana
- `grafana-dashboard.json`: Preconfigured Grafana dashboard
- `provisioning/`: Auto-loads Prometheus data source and dashboard into Grafana

---

## Setup Go Project

```bash
go mod tidy
go run main.go
```

The server starts on `http://localhost:8080`.

---

## What `main.go` Does

- Exposes `/ping` for a test endpoint
- Exposes `/metrics` for Prometheus to scrape
- Uses Prometheus `Counter` and `Histogram`:
  - `http_requests_total`
  - `http_request_duration_seconds`

---

## 🐳 Run Prometheus + Grafana via Docker

```bash
docker compose up
```

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (login: `admin` / `admin`)

---

## 📊 Default Dashboard (Grafana)

The `grafana-dashboard.json` shows:

- **Requests per Second** – real-time request throughput
- **95th Percentile Latency** – tail latency using `histogram_quantile(0.95, ...)`
- **Average Duration** – derived from histogram sum/count

---

## 📈 How to Use Prometheus + Grafana

1. Open Prometheus: http://localhost:9090
   - Try queries like:
     - `http_requests_total`
     - `rate(http_requests_total[1m])`
     - `histogram_quantile(0.95, sum(rate(http_request_duration_seconds_bucket[5m])) by (le))`

2. Open Grafana: http://localhost:3000
   - Dashboards → “Go App Metrics”
   - Visualize trends over time
   - Edit panels to explore PromQL

---

## 📚 Prometheus + PromQL Learning Resources

- [Prometheus Docs](https://prometheus.io/docs/introduction/overview/)
- [PromQL Tutorial (official)](https://prometheus.io/docs/prometheus/latest/querying/basics/)
- [Grafana Prometheus Queries](https://grafana.com/docs/grafana/latest/datasources/prometheus/)
- [Monitoring System Design](https://sre.google/sre-book/monitoring-distributed-systems/)
