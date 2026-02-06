package profile

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/AY88o/switchblade/internal/sys"
)

type Profile struct {
	Name string
	Apps []string
}

func (p Profile) Start() {
	fmt.Printf("\n--- IGNITING %s PROTOCOL ---\n", p.Name)

	for _, app := range p.Apps {

		fmt.Printf("[+] Launching %s...\n", app)
		cmdName := app
		if len(app) > 4 && app[len(app)-4:] == ".exe" {
			cmdName = app[:len(app)-4]
		}
		cmd := exec.Command("cmd", "/C", "start", "", cmdName)
		cmd.Start()
		time.Sleep(500 * time.Millisecond)
	}
}

func Kill(list []string) {

	fmt.Printf("...Closing Current state...\n")

	for _, app := range list {

		cmd := exec.Command("taskkill", "/IM", app, "/F", "/T")

		cmd.Run()

	}

	fmt.Printf("Kill succesful, closed %d apps\n", len(list))

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
