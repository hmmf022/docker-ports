package output

import (
	"bytes"
	"strings"
	"testing"

	"docker-ports/internal/docker"
)

func TestTable(t *testing.T) {
	t.Parallel()
	rows := []docker.ContainerPorts{
		{Name: "api", LocalPorts: []int{8080}},
		{Name: "db", LocalPorts: []int{5432, 15432}},
	}

	var b bytes.Buffer
	if err := Table(&b, rows); err != nil {
		t.Fatalf("Table() error = %v", err)
	}

	out := b.String()
	for _, expected := range []string{"NAME", "LOCAL_PORTS", "api", "8080", "db", "5432,15432"} {
		if !strings.Contains(out, expected) {
			t.Fatalf("output missing %q:\n%s", expected, out)
		}
	}
}
