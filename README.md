# BIG-IP Exporter

[![Go](https://img.shields.io/badge/Go-1.23-blue.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A **Prometheus exporter** for F5 BIG-IP devices using the iControl REST API.

It follows the standard blackbox exporter pattern (`/probe?target=...`) and supports monitoring multiple BIG-IP devices concurrently.

## Features

- Virtual Servers (status, traffic, connections, etc.)
- Pools and Pool Members statistics
- System / Hardware resources
- Disk usage
- SSL Certificates (expiration, status)
- HA / Sync-Failover status
- Token-based authentication
- Concurrent multi-target scraping

## Quick Start

### 1. Build

```bash
git clone https://github.com/Haameed/bigip_exporter.git
cd bigip_exporter
go build -o bigip_exporter .
```

### 2. Configuration

Create a file named `bigip-config.yaml`:

```yaml
targets:
  "https://192.168.1.10":
    username: "admin"
    password: "StrongPassword123"
  "https://192.168.1.11":
    username: "admin"
    password: "StrongPassword123"
```

### 3. Run

```bash
./bigip_exporter -config bigip-config.yaml
```

### 4. Prometheus Scrape Config

```yaml
scrape_configs:
  - job_name: 'bigip'
    metrics_path: '/probe'
    params:
      target: ['https://192.168.1.10']
    static_configs:
      - targets: ['localhost:9711']
```

## Command Line Flags

| Flag                    | Default                | Description                        |
|-------------------------|------------------------|------------------------------------|
| `-config`               | `bigip-config.yaml`    | Path to config file                |
| `-listen`               | `:9711`                | HTTP listen address                |
| `-scrape-timeout`       | `30`                   | Scrape timeout in seconds          |

## Endpoints

- `GET /probe?target=https://<bigip-ip>` → Returns metrics for one target
- `GET /metrics` → Exporter internal metrics
- `GET /health` → Health check

## Contributing

**Everyone is welcome to participate and contribute!** 🎉

Feel free to open issues for bugs or new feature requests, and submit Pull Requests.

## License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.
