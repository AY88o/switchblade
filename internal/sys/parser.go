package sys

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Capture() ([]string, error) {

	psCmd := `Get-Process | Where-Object {$_.MainWindowTitle -ne "" -and $_.Path -ne $null} | Select-Object -ExpandProperty Path`

	cmd := exec.Command("powershell", "-Command", psCmd)
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("powershell Error %v", err)
		return nil, err
	}

	lines := strings.Split(string(output), "\r\n")

	seenAppPaths := make(map[string]bool)

	for _, path := range lines {
		//powershell may return whitespaces
		path = strings.TrimSpace(path)
		path = strings.ToLower(path)

		//to not capture the switchblade or powershell/terminal itself
		appName := filepath.Base(path)

		//path.Trimspace is strict
		if path == "" {
			continue
		}

		if appName == "switchblade.exe" ||
			appName == "main.exe" ||
			appName == "WindowsTerminal.exe" ||
			appName == "cmd.exe" ||
			appName == "explorer.exe" ||
			appName == "ApplicationFrameHost.exe" ||
			appName == "SearchHost.exe" ||
			appName == "StartMenuExperienceHost.exe" {
			continue
		}

		seenAppPaths[path] = true

	}

	cleanListPaths := make([]string, 0, len(seenAppPaths))

	for paths := range seenAppPaths {
		cleanListPaths = append(cleanListPaths, paths)
	}

	return cleanListPaths, nil

}

func Subtract(mixList []string, pureNoiseList []string) []string {

	noiseMap := make(map[string]bool)

	for _, app := range pureNoiseList {
		noiseMap[app] = true
	}

	var cleanList []string

	for _, app := range mixList {

		if !noiseMap[app] {
			cleanList = append(cleanList, app)
		}
	}

	return cleanList
}
