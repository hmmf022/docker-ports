package docker

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	hostPortPattern  = regexp.MustCompile(`(?:\[::\]|[0-9.]+):([0-9]+)->`)
	errDockerMissing = errors.New("docker CLI not found; install Docker and ensure `docker` is in PATH")
)

type psRow struct {
	Names string `json:"Names"`
	Ports string `json:"Ports"`
}

type ContainerPorts struct {
	Name       string `json:"name"`
	LocalPorts []int  `json:"local_ports"`
}

func ListPublishedPorts(all bool) ([]ContainerPorts, error) {
	args := []string{"ps"}
	if all {
		args = append(args, "-a")
	}
	args = append(args, "--format", "{{json .}}")

	cmd := exec.Command("docker", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return nil, errDockerMissing
		}
		return nil, fmt.Errorf("failed to run `docker %s`: %w\n%s", strings.Join(args, " "), err, strings.TrimSpace(string(out)))
	}

	results := make([]ContainerPorts, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var row psRow
		if err := json.Unmarshal([]byte(line), &row); err != nil {
			return nil, fmt.Errorf("failed to parse docker output line %q: %w", line, err)
		}

		ports := extractPublishedHostPorts(row.Ports)
		if len(ports) == 0 {
			continue
		}

		results = append(results, ContainerPorts{
			Name:       row.Names,
			LocalPorts: ports,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading docker output: %w", err)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})
	return results, nil
}

func extractPublishedHostPorts(portsField string) []int {
	matches := hostPortPattern.FindAllStringSubmatch(portsField, -1)
	if len(matches) == 0 {
		return nil
	}

	seen := make(map[int]struct{}, len(matches))
	ports := make([]int, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		p, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}
		ports = append(ports, p)
	}

	sort.Ints(ports)
	return ports
}
