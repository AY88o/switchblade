package profile

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/AY88o/switchblade/internal/sys"
)

type Profile struct {
	Name string
	Apps []string
}

func (p Profile) Start() {
	fmt.Printf("\n--- IGNITING %s PROTOCOL ---\n", p.Name)

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

	fmt.Printf("...Closing Current state...\n")

	for _, appPath := range list {

		appName := filepath.Base(appPath)

		if appName == "switchblade.exe" || appName == "main.exe" {
			continue
		}

		cmd := exec.Command("taskkill", "/IM", appName, "/F", "/T")
		err := cmd.Run()

		if err != nil {
			fmt.Printf("Error Killing the state %v", err)
		}

	}

}

func CloseCurrentState() error {

	fmt.Println("killing the Current state...")
	//Capturing current state
	fmt.Println("Capturing only to filter...")
	CurrentStateList, err := sys.Capture()

	//summoning prev noise to filter
	pureNoiseListStruct, err := LoadProfile("Noise")

	if err != nil {
		fmt.Println("Error, couldnt find the calibration file")
		return err
	}

	//filtering a clear list
	fmt.Println("filtering to kill the state...")
	clearList := sys.Subtract(CurrentStateList, pureNoiseListStruct.Apps)

	//killing the state
	Kill(clearList)
	fmt.Println("State Killed Successfully...")

	return nil

}

func OpenSavedState(stateName string) error {

	fmt.Printf("Loading the saved state %s ...\n", stateName)
	savedProfileStruct, err := LoadProfile(stateName)

	if err != nil {
		fmt.Printf("Error loading Saved profile, Saved profile doesnt exits: %v\n", err)
		return err
	}

	savedProfileStruct.Start()
	fmt.Println("Success loading the state!")
	return nil

}
