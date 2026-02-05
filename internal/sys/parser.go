package sys

import (
	"encoding/csv"
	"fmt"
	"os/exec"
	"strings"
)

func Capture() ([]string, error) {

	cmd := exec.Command("tasklist", "/V", "/FO", "CSV", "/NH")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running tasklist:", err)
		return nil, fmt.Errorf("Error running tasklist: %w", err)
	}

	reader := csv.NewReader(strings.NewReader(string(output)))

	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed to parse tasklist output: %w", err)
	}

	seenApps := make(map[string]bool)

	for _, row := range records {
		if len(row) < 9 {
			continue
		}
		appName := row[0]
		windowsTitle := row[8]

		if appName == "explorer.exe" || appName == "tasklist.exe" || appName == "svchost.exe" {
			continue
		}

		if !strings.HasSuffix(appName, ".exe") {
			continue
		}

		if windowsTitle == "N/A" {
			continue
		}

		seenApps[appName] = true
	}

	cleanList := make([]string, 0, len(seenApps))
	// var cleanList []string

	for unique := range seenApps {
		cleanList = append(cleanList, unique)
	}

	return cleanList, nil
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
