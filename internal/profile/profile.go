package profile

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/AY88o/switchblade/internal/sys"
)

type Profile struct {
	Name string
	Apps []string
}

func (p Profile) Start() {

	for _, appPath := range p.Apps {

		cmd := exec.Command("cmd", "/C", "start", "", appPath)
		err := cmd.Start()

		if err != nil {
			fmt.Printf("Error starting the state %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func Kill(list []string) {
	fmt.Println("")

	fmt.Printf("CLOSING THE CURRENT STATE...\n")

	fmt.Println("")

	for _, appPath := range list {

		appName := filepath.Base(appPath)

		if appName == "switchblade.exe" || appName == "main.exe" {
			continue
		}

		cmd := exec.Command("taskkill", "/IM", appName, "/F", "/T")
		err := cmd.Run()

		if err != nil {
			fmt.Printf("Error Killing the state %v\n", err)
		}

	}

}

func CloseCurrentState(force bool, interactive bool) error {

	CurrentStateList, err := sys.Capture()

	//summoning prev noise to filter
	pureNoiseListStruct, err := LoadProfile("Noise")

	if err != nil {
		fmt.Println("Error, couldnt find the calibration file")
		return err
	}

	clearList := sys.Subtract(CurrentStateList, pureNoiseListStruct.Apps)

	if force {
		//killing the state
		Kill(clearList)
		fmt.Println("State Killed Successfully...")

	} else if interactive {
		fmt.Println("")
		fmt.Printf("  %-25s\n", "INTERACTIVE KILL MODE")
		fmt.Println("-----------------------")
		fmt.Println("The following apps will be terminated:")

		for _, app := range clearList {
			app = filepath.Base(app)
			app = strings.TrimSuffix(app, ".exe")
			fmt.Printf("  - %s\n", app)
		}

		fmt.Println("")

		var Permission string
		fmt.Println("")
		fmt.Println("Are you sure you want to continue?")
		fmt.Print(" [Y] Yes, kill them   [N] No, abort > ")

		fmt.Scan(&Permission)

		if Permission == "y" || Permission == "Y" {
			//killing the state
			Kill(clearList)
			fmt.Println("")
			fmt.Println("State Killed Successfully...")
			fmt.Println("")

		} else {
			fmt.Println("")
			fmt.Println("Kill aborted...")
			fmt.Println("")
		}
	}

	return nil

}

func OpenSavedState(stateName string) error {

	fmt.Println("")

	fmt.Printf("LOADING THE SAVED STATE %s\n", stateName)
	fmt.Println("")
	savedProfileStruct, err := LoadProfile(stateName)

	if err != nil {
		fmt.Printf("Error loading Saved profile, Saved profile doesnt exits: %v\n", err)
		return err
	}

	savedProfileStruct.Start()

	fmt.Println("Success loading the state!")
	fmt.Println("")
	return nil

}
