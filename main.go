package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type LoadState string
type ActiveState string
type SubState string

const (
	LoadStateLoaded   LoadState = "loaded"
	LoadStateNotFound LoadState = "not-found"
)

const (
	ActiveStateActive   ActiveState = "active"
	ActiveStateInactive ActiveState = "inactive"
	ActiveStateFailed   ActiveState = "failed"
)

const (
	SubStateRunning SubState = "running"
	SubStateDead    SubState = "dead"
	SubStateFailed  SubState = "failed"
)

type ServiceStatus struct {
	Name        string
	LoadState   LoadState
	ActiveState ActiveState
	SubState    SubState
}

func (s ServiceStatus) Exists() bool {
	return s.LoadState == LoadStateLoaded
}

func (s ServiceStatus) IsRunning() bool {
	return s.ActiveState == ActiveStateActive &&
		s.SubState == SubStateRunning
}

func main() {
	serviceName := "nginx"

	cmd := exec.Command(
		"systemctl",
		"show",
		serviceName,
		"--property=LoadState",
		"--property=ActiveState",
		"--property=SubState",
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("failed to inspect service: %v\n", err)
		return
	}

	status := ServiceStatus{
		Name: serviceName,
	}

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		key, value, found := strings.Cut(line, "=")

		if !found {
			continue
		}

		switch key {
		case "LoadState":
			status.LoadState = LoadState(value)

		case "ActiveState":
			status.ActiveState = ActiveState(value)

		case "SubState":
			status.SubState = SubState(value)
		}
	}

	fmt.Printf("Service: %s\n", status.Name)
	fmt.Printf("Load State: %s\n", status.LoadState)
	fmt.Printf("Active State: %s\n", status.ActiveState)
	fmt.Printf("Sub State: %s\n", status.SubState)
	fmt.Printf("Exists: %t\n", status.Exists())
	fmt.Printf("Running: %t\n", status.IsRunning())
}
