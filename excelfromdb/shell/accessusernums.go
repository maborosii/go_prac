package shell

import (
	_ "embed"
	"os/exec"
)

//go:embed Statistics.sh
var accessstat string

func AccessStatOutput() (string, error) {
	cmd := exec.Command("/bin/bash", "-c", accessstat)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
