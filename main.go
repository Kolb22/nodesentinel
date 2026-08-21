package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"systemctl",
		"is-active",
		"nginx",
	)

	output, err := cmd.CombinedOutput()

	fmt.Printf("Output: %s", output)

	if err != nil {
		fmt.Printf("Error type: %T\n", err)
		fmt.Printf("Error: %v\n", err)
	}
}
