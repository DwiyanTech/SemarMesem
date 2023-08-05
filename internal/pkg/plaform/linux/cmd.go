package cmd

import "os/exec"

func shellOutput(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()
	out := string(output)
	return out, err
}