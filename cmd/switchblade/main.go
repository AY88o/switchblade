package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AY88o/switchblade/internal/profile"
	"github.com/AY88o/switchblade/internal/sys"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Something went wrong!")
		fmt.Printf("Error: %v\n", r)
		fmt.Println("Check command syntax and try again.")
	}
}

func main() {

	defer handlePanic()

	if len(os.Args) < 2 {

		printHelp()

		return

	}

	command := os.Args[1]

	switch command {

	case "calibrate":
		runCalibrate()

	case "save":
		runSave()

	case "go":
		runGo()

	case "help":
		printHelp()

	default:
		fmt.Printf("Uknown command: %s\n", command)
		printHelp()
	}

}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("switchblade calibrate     ...(calibrate the tool)")
	fmt.Println("OR")
	fmt.Println("switchblade save <name>   ...........(Save state)")
	fmt.Println("OR")
	fmt.Println("switchblade go <name>      ...(open a saved state)")
	fmt.Println("OR")
	fmt.Println("switchblade go -k <name>   ...(Permission to kill, open saved stat)")
	fmt.Println("switchblade go -fk <name>   ....(force kill current state, open saved state)")
}

func runCalibrate() {
	fmt.Println("Scanning system for noise...")

	pureNoiseList, err := sys.Capture()
	fmt.Printf("Caught %d ghost apps in the background\n", len(pureNoiseList))

	if err != nil {
		fmt.Printf("Error Capturing: %v\n", err)
	}

	baseProfile := profile.Profile{

		Name: "Noise",
		Apps: pureNoiseList,
	}

	fmt.Println("Recognising noise for future analysis")

	err = profile.SaveProfile(baseProfile)
	if err != nil {
		fmt.Printf("Erorr Saving profile: %v\n", err)
	}

	fmt.Println("Calibration Done!")
}

func runSave() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("Give the state a name...")
		fmt.Println("switchblade save <name>")

		return
	}

	profilname := os.Args[2]

	fmt.Println("Capturing...")

	mixList, err := sys.Capture()

	if err != nil {
		fmt.Printf("Error Capturing: %v\n", err)
	}

	fmt.Println("Filtering...")

	pureNoiseListStruct, err := profile.LoadProfile("Noise")

	if err != nil {
		fmt.Println("Error, couldnt find the calibration file")
	}

	fmt.Println("Saving state...")
	clearList := sys.Subtract(mixList, pureNoiseListStruct.Apps)

	Aprofile := profile.Profile{
		Name: profilname,
		Apps: clearList,
	}

	err = profile.SaveProfile(Aprofile)
	if err != nil {
		fmt.Printf("Error Saving the state: %v", err)
	}

	fmt.Printf("State Name: %s\n", profilname)
	for _, app := range clearList {
		app = filepath.Base(app)
		fmt.Println(app)
	}

	fmt.Printf("State Saved Successfully!")
}

func runGoHelp() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("switchblade go <name>  .....(open already saved state)")
		fmt.Println("switchblade go -k <name> ...(open saved state and kill current state)")
	}
}

func runKillandSwitch(force bool, interactive bool) {
	if len(os.Args) < 4 {
		fmt.Println("Usage:")
		fmt.Println("switchblade go -k <name>  ...(open saved state and kill current one)")
		return

	}

	//killing current state
	err := profile.CloseCurrentState(force, interactive)

	if err != nil {
		fmt.Printf("couldn't kill the current state : %v", err)
		return
	}

	//opening the saved state
	savedProfileName := os.Args[3]

	err2 := profile.OpenSavedState(savedProfileName)

	if err2 != nil {
		fmt.Printf("Error opening saved state %s", savedProfileName)
		return
	}
}

func runOpen(arg string) {
	savedProfileName := arg

	err := profile.OpenSavedState(savedProfileName)

	if err != nil {
		fmt.Printf("Error opening saved state %s", savedProfileName)
		return
	}
}

func runGo() {

	if len(os.Args) < 3 {
		runGoHelp()
		return
	}

	flag := os.Args[2]

	switch flag {
	case "-fk":
		runKillandSwitch(true, false)

	case "-k":
		runKillandSwitch(false, true)

	default:
		runOpen(flag)
	}

}
