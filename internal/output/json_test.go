package output

import (
	"bytes"
	"encoding/json"
	"testing"

	"docker-ports/internal/docker"
)

func TestJSON(t *testing.T) {
	t.Parallel()
	rows := []docker.ContainerPorts{{Name: "web", LocalPorts: []int{8080, 8443}}}

	var b bytes.Buffer
	if err := JSON(&b, rows); err != nil {
		t.Fatalf("JSON() error = %v", err)
	}

	var got []docker.ContainerPorts
	if err := json.Unmarshal(b.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal JSON output: %v", err)
	}

	if len(got) != 1 || got[0].Name != "web" || len(got[0].LocalPorts) != 2 {
		t.Fatalf("unexpected JSON payload: %+v", got)
	}
}
