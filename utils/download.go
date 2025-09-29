package utils

import "fmt"

const DOWNLOAD_URL = "https://go.dev/dl"

func GetAliasArch(arch string) string {
	switch arch {
	case "x86_64", "x86-64":
		{
			return "amd64"
		}
	case "ARM64":
		{
			return "arm64"
		}
	case "x86":
		{
			return "386"
		}
	default:
		{
			return ""
		}
	}
}

func GetAliasOs(os string) string {
	switch os {
	case "Linux":
		{
			return "linux"
		}
	case "Windows":
		{
			return "windows"
		}
	case "macOS":
		{
			return "darwin"
		}
	default:
		{
			return ""
		}
	}
}

func GetFileSuffix(os, arch string) string {
	aliasArch := GetAliasArch(arch)
	aliasOs := GetAliasOs(os)
	if aliasArch == "" && aliasOs == "" {
		return "src"
	}

	return fmt.Sprintf("%s-%s", aliasOs, aliasArch)
}

func GetFilename(os, arch, version string) string {
	suffix := GetFileSuffix(os, arch)
	return fmt.Sprintf("%s.%s.tar.gz", version, suffix)
}

func GetDownloadPath(os, arch, version string) string {
	filename := GetFilename(os, arch, version)
	return fmt.Sprintf("%s/%s", DOWNLOAD_URL, filename)
}
