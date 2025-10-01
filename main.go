package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alecthomas/kong"
	"github.com/nguyenluan2001/gov/controller"
	"github.com/nguyenluan2001/gov/utils"
	// "github.com/nguyenluan2001/gov/controller"
)

func main() {
	var CLI struct {
		Install struct {
			Version string `arg:"" name:"version" help:"Go version to install."`
		} `cmd:"" help:"Install specific go version."`
		Use struct {
			Version string `arg:"" name:"version" help:"Select go version."`
		} `cmd:"" help:"Select go version."`
		List struct {
		} `cmd:"" help:"Install specific go version."`
	}
	cwd, _ := os.Getwd()
	app := controller.App{
		TempPath: path.Join(cwd, ".gov"),
		RootPath: cwd,
		HomePath: "/root",
	}
	ctx := kong.Parse(&CLI, kong.Name("gov"))
	switch ctx.Command() {
	case "install <version>":
		{
			app.InstallCmd(CLI.Install.Version)

			bashrcPath := path.Join(app.HomePath, ".bashrc")
			utils.UpdateBashrc(fmt.Sprintf("export PATH=%s:$PATH\n", path.Join(app.TempPath, "current", "bin")), bashrcPath)
		}
	case "use <version>":
		{
			app.UseCmd(CLI.Use.Version)
		}
	case "list":
		{
			app.ListCmd()
		}
	default:
		panic(ctx.Command())
	}
}
