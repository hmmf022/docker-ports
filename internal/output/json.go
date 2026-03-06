package output

import (
	"encoding/json"
	"io"

	"docker-ports/internal/docker"
)

func JSON(w io.Writer, rows []docker.ContainerPorts) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(rows)
}
