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
- `--json`: Output in JSON format

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

[License](./LICENSE)
