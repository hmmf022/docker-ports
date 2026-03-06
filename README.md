[English README](./README.en.md)

# docker-ports

`docker ps` から、`docker exec` で使うコンテナ名と localhost からアクセスできる公開ポートを一覧表示するCLIです。

## Requirements

- Go 1.22+
- Docker CLI (`docker`) が利用可能であること

## Usage

### Run directly

```bash
go run ./cmd/docker-ports
```

### Install as command

```bash
go install ./cmd/docker-ports
```

インストール後:

```bash
docker-ports
```

## Options

- `--all`: 停止中コンテナも含めて表示
- `--json`: JSON形式で出力
- `--version`: バージョン情報を表示

## Examples

```bash
docker-ports
```

出力例:

```text
NAME   LOCAL_PORTS
api    8080
db     5432,15432
```

```bash
docker-ports --json
```

出力例:

```json
[
  {
    "name": "api",
    "local_ports": [8080]
  }
]
```

## Install from Releases

GitHub Releases の Assets から自分の環境向けアーカイブをダウンロードして展開してください。

- Linux/macOS: `docker-ports_<version>_<os>_<arch>.tar.gz`
- Windows: `docker-ports_<version>_<os>_<arch>.zip`
- チェックサム: `checksums.txt` (SHA256)

[License](./LICENSE)
