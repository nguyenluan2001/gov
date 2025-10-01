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

func UnTarFile(backUrl, extractToUrl, folderName, path string) error {
	os.Chdir(extractToUrl)
	os.Mkdir(folderName, 0777)
	cmd := exec.Command("tar", "-xvzf", path, "--directory", folderName, "--strip-components=1")
	output, err := cmd.Output()
	fmt.Println("UnTarFile", string(output), backUrl, extractToUrl)
	time.Sleep(5 * time.Second)
	os.Chdir(backUrl)
	return err
}

func CreateSymbolLink(source, target string) {
	//ln -s /app/.gov/go1.24.7 /app/.gov/current
	unlinkCmd := exec.Command("unlink", target)
	unlinkCmd.Output()
	cmd := exec.Command("ln", "-s", source, target)
	cmd.Output()
}

func SourceFile(path string) {
	cmd := exec.Command("source", path)
	cmd.Output()
}

func UpdateBashrc(content, bashrcPath string) {
	file, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalln("Open file failed.", err)
	}
	file.WriteString(content)
	defer file.Close()
	SourceFile(bashrcPath)
}
