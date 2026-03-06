package output

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"

	"docker-ports/internal/docker"
)

func Table(w io.Writer, rows []docker.ContainerPorts) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	if _, err := fmt.Fprintln(tw, "NAME\tLOCAL_PORTS"); err != nil {
		return err
	}
	for _, row := range rows {
		if _, err := fmt.Fprintf(tw, "%s\t%s\n", row.Name, joinPorts(row.LocalPorts)); err != nil {
			return err
		}
	}
	return tw.Flush()
}

func joinPorts(ports []int) string {
	if len(ports) == 0 {
		return ""
	}
	parts := make([]string, 0, len(ports))
	for _, p := range ports {
		parts = append(parts, strconv.Itoa(p))
	}
	return strings.Join(parts, ",")
}
