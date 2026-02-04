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

}
