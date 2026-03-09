[日本語 README](./README.md)

# docker-ports

CLI tool that lists container names (for `docker exec`) and localhost-accessible published ports from `docker ps`.

## Requirements

- Go 1.22+
- Docker CLI (`docker`) must be available

## Usage

### Run directly

```bash
go run ./cmd/docker-ports
```

### Install as command

```bash
go install ./cmd/docker-ports
```

After installation:

```bash
docker-ports
```

## Options

- `--all`: Include stopped containers
- `--name`: Filter by container name (case-insensitive substring)
- `--json`: Output in JSON format
- `--version`: Print version information

## Examples

```bash
docker-ports
```

Example output:

```text
NAME   LOCAL_PORTS
api    8080
db     5432,15432
```

```bash
docker-ports --name api
```

Example output:

```text
NAME          LOCAL_PORTS
api           8080
my-api-worker 18080
```

```bash
docker-ports --json
```

Example output:

```json
[
  {
    "name": "api",
    "local_ports": [8080]
  }
]
```

## Install from Releases

Download and extract the archive for your platform from GitHub Releases assets.

- Linux/macOS: `docker-ports_<version>_<os>_<arch>.tar.gz`
- Windows: `docker-ports_<version>_<os>_<arch>.zip`
- Checksums: `checksums.txt` (SHA256)

[License](./LICENSE)
