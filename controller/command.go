package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/nguyenluan2001/gov/utils"
)

type App struct {
	TempPath string `json:"tempPath"`
	RootPath string `json:"rootPath"`
	HomePath string `json:"homePath"`
}

func (app *App) InstallCmd(version string) {
	fileInfo, err := os.Stat(path.Join(app.TempPath, version))
	if err == nil && fileInfo.IsDir() {
		fmt.Println("This go version has already installed.")
		return
	}
	osName := utils.GetOs()
	arch := utils.GetArch()
	fmt.Println(osName, arch)

	//Download file
	filename := utils.GetFilename(strings.TrimSpace(osName), strings.TrimSpace(arch), version)
	downloadPath := utils.GetDownloadPath(strings.TrimSpace(osName), strings.TrimSpace(arch), version)
	fmt.Println("downloadPath", downloadPath)
	filepath := path.Join(app.TempPath, filename)
	file, err := os.Create(filepath)

	if err != nil {
		log.Fatalln("Create file failed.")
	}

	resp, respErr := http.Get(downloadPath)

	if respErr != nil {
		log.Fatalln("Download file failed")
	}
	io.Copy(file, resp.Body)
	extractErr := utils.UnTarFile(app.RootPath, app.TempPath, version, filepath)
	fmt.Println("extractErr", extractErr)
	if extractErr != nil {
		log.Fatalln("Download file failed")
	}

	sourceUrl := path.Join(app.TempPath, version)
	targetUrl := path.Join(app.TempPath, "current")
	utils.CreateSymbolLink(sourceUrl, targetUrl)
	os.Remove(filepath)
}

func (app *App) UseCmd(version string) {
	sourceUrl := path.Join(app.TempPath, version)
	targetUrl := path.Join(app.TempPath, "current")
	utils.CreateSymbolLink(sourceUrl, targetUrl)
}

func (app *App) UnInstallCmd() {
}

func (app *App) ListCmd() {
	entries, _ := os.ReadDir(app.TempPath)
	versions := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		versions = append(versions, entry.Name())
	}
	text := fmt.Sprintf("Installed versions:\n%s", strings.Join(versions, "\n"))
	fmt.Println(text)
}
