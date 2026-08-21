package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type ServiceStatus string

const (
	StatusActive   ServiceStatus = "active"
	StatusInactive ServiceStatus = "inactive"
	StatusFailed   ServiceStatus = "failed"
	StatusUnknown  ServiceStatus = "unknown"
)

func parseServiceStatus(output string) ServiceStatus {
	switch output {
	case "active":
		return StatusActive
	case "inactive":
		return StatusInactive
	case "failed":
		return StatusFailed
	default:
		return StatusUnknown
	}
}

func main() {
	cmd := exec.Command(
		"systemctl",
		"is-active",
		"nginx",
	)

	output, err := cmd.CombinedOutput()

	rawStatus := strings.TrimSpace(string(output))
	status := parseServiceStatus(rawStatus)

	fmt.Printf("Service: nginx\n")
	fmt.Printf("Status: %s\n", status)

	if err != nil {
		exitError, ok := err.(*exec.ExitError)

		if ok {
			fmt.Printf(
				"Exit Code: %d\n",
				exitError.ExitCode(),
			)
		} else {
			fmt.Printf(
				"Unexpected error: %v\n",
				err,
			)
		}
	}
}
