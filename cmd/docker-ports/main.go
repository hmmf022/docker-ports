package main

import (
	"flag"
	"fmt"
	"os"

	"docker-ports/internal/docker"
	"docker-ports/internal/output"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	showAll := flag.Bool("all", false, "Include stopped containers")
	asJSON := flag.Bool("json", false, "Output as JSON")
	showVersion := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("docker-ports version=%s commit=%s date=%s\n", version, commit, date)
		return
	}

	containers, err := docker.ListPublishedPorts(*showAll)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *asJSON {
		err = output.JSON(os.Stdout, containers)
	} else {
		err = output.Table(os.Stdout, containers)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "output error: %v\n", err)
		os.Exit(1)
	}
}
