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
		fmt.Println("switchblade calibrate  ....(calibrate the tool)")
		fmt.Println("OR")
		fmt.Println("switchblade save <name>   ...(Save state)")

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

		pureNoiseList, err := profile.LoadProfile("Noise")

		if err != nil {
			fmt.Println("Error, couldnt find the calibration file")
		}

		clearList := sys.Subtract(mixList, pureNoiseList)

		Aprofile := profile.Profile{
			Name: profilname,
			Apps: clearList,
		}

		err = profile.SaveProfile(Aprofile)

	}

	// if command == "go" {
	// 	if len(os.Args) < 3 {
	// 		fmt.Println("Usage name:")
	// 		fmt.Println("Enter the Saved state name")
	// 		fmt.Println("switchblade go <name>")

	// 		return
	// 	}

	// 	stateName := os.Args[3]

	// }

}
