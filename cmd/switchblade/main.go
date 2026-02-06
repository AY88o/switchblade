package main

import (
	"fmt"
	"os"

	"github.com/AY88o/switchblade/internal/profile"
	"github.com/AY88o/switchblade/internal/sys"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println("Usage:")
		fmt.Println("switchblade calibrate     ...(calibrate the tool)")
		fmt.Println("OR")
		fmt.Println("switchblade save <name>   ...........(Save state)")
		fmt.Printf("switchblade go <name>      ...(open a saved state)")
		fmt.Printf("switchblade go -k <name>   ...(kill the current state and open a saved state)")

		return

	}

	command := os.Args[2]

	if command == "calibrate" {

		fmt.Println("Scanning system for noise...")

		pureNoiseList, err := sys.Capture()

		if err != nil {
			fmt.Printf("Error Capturing: %v\n", err)
		}

		baseProfile := profile.Profile{

			Name: "Noise",
			Apps: pureNoiseList,
		}

		err = profile.SaveProfile(baseProfile)
		if err != nil {
			fmt.Printf("Erorr Saving profile: %v\n", err)
		}

	}

	if command == "save" {
		if len(os.Args) < 4 {
			fmt.Println("Usage name:")
			fmt.Println("Give the state a name:")
			fmt.Println("switchblade save <name>")

			return
		}

		profilname := os.Args[3]

		fmt.Println("Capturing...")

		mixList, err := sys.Capture()

		if err != nil {
			fmt.Printf("Error Capturing: %v\n", err)
		}

		pureNoiseListStruct, err := profile.LoadProfile("Noise")

		if err != nil {
			fmt.Println("Error, couldnt find the calibration file")
		}

		clearList := sys.Subtract(mixList, pureNoiseListStruct.Apps)

		Aprofile := profile.Profile{
			Name: profilname,
			Apps: clearList,
		}

		err = profile.SaveProfile(Aprofile)

	}

	if command == "go" {
		if len(os.Args) < 3 {
			fmt.Println("Usage:")
			fmt.Println("switchblade go <name>  .....(open already saved state)")
			fmt.Println("switchblade go -k <name> ...(open saved state and kill current state)")
			return
		}

		if os.Args[3] == "-k" {
			if len(os.Args) < 4 {
				fmt.Println("Usage:")
				fmt.Println("switchblade go -k <name>  ...(open saved state and kill current one)")
				return

			}

			//killing current state

			err := profile.CloseCurrentState()

			if err != nil {
				fmt.Printf("couldn't kill the current state : %v", err)
				return
			}

			//opening the saved state
			savedProfileName := os.Args[4]

			err2 := profile.OpenSavedState(savedProfileName)

			if err2 != nil {
				fmt.Printf("Error opening saved state %s", savedProfileName)
				return
			}

		} else {
			savedProfileName := os.Args[3]

			err := profile.OpenSavedState(savedProfileName)

			if err != nil {
				fmt.Printf("Error opening saved state %s", savedProfileName)
				return
			}
		}

	}

}
