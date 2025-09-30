package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
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

func UnTarFile(backUrl, extractToUrl, path string) error {
	os.Chdir(extractToUrl)
	cmd := exec.Command("tar", "-xvzf", path)
	output, err := cmd.Output()
	fmt.Println("UnTarFile", string(output), backUrl, extractToUrl)
	time.Sleep(5 * time.Second)
	os.Chdir(backUrl)
	return err
}
