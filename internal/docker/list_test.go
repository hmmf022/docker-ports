package docker

import (
	"reflect"
	"testing"
)

func TestExtractPublishedHostPorts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want []int
	}{
		{
			name: "empty",
			in:   "",
			want: nil,
		},
		{
			name: "non published only",
			in:   "80/tcp, 443/tcp",
			want: nil,
		},
		{
			name: "single published",
			in:   "0.0.0.0:8080->80/tcp",
			want: []int{8080},
		},
		{
			name: "ipv4 and ipv6 duplicate",
			in:   "0.0.0.0:8080->80/tcp, [::]:8080->80/tcp",
			want: []int{8080},
		},
		{
			name: "multiple unique",
			in:   "0.0.0.0:8080->80/tcp, [::]:8080->80/tcp, 127.0.0.1:5432->5432/tcp, 0.0.0.0:8443->443/tcp",
			want: []int{5432, 8080, 8443},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := extractPublishedHostPorts(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("extractPublishedHostPorts(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestFilterContainersByName(t *testing.T) {
	t.Parallel()

	rows := []ContainerPorts{
		{Name: "api", LocalPorts: []int{8080}},
		{Name: "my-API", LocalPorts: []int{18080}},
		{Name: "db", LocalPorts: []int{5432}},
	}

	tests := []struct {
		name  string
		query string
		want  []ContainerPorts
	}{
		{
			name:  "empty query returns all",
			query: "",
			want:  rows,
		},
		{
			name:  "substring case insensitive",
			query: "api",
			want: []ContainerPorts{
				{Name: "api", LocalPorts: []int{8080}},
				{Name: "my-API", LocalPorts: []int{18080}},
			},
		},
		{
			name:  "no matches",
			query: "cache",
			want:  []ContainerPorts{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := filterContainersByName(rows, tt.query)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("filterContainersByName(%v, %q) = %v, want %v", rows, tt.query, got, tt.want)
			}
		})
	}
}
