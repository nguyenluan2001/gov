package utils

import (
	"log"
	"os/exec"
)

func GetOs() string {
	cmd := exec.Command("uname", "-s")
	os, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	// Print the output
	return string(os)
}

func GetArch() string {
	cmd := exec.Command("uname", "-m")
	arch, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	// Print the output
	return string(arch)
}

func UnTarFile(path string) error {
	cmd := exec.Command("tar", "-xvzf", path)
	_, err := cmd.Output()
	return err
}
