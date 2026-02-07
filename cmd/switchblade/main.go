package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	case "list", "ls":
		runList()

	default:
		fmt.Printf("Uknown command: %s\n", command)
		printHelp()
	}

}

func printHelp() {
	fmt.Println("\nSWITCHBLADE v1.3")
	fmt.Println("----------------")
	fmt.Println("Usage:")
	fmt.Println("  switchblade <command> [arguments]")
	fmt.Println("")

	fmt.Println("Core Commands:")

	fmt.Printf("  %-25s %s\n", "calibrate", "Scan system to learn background noise")
	fmt.Printf("  %-25s %s\n", "save <name>", "Save currently running apps as a profile")
	fmt.Printf("  %-25s %s\n", "ls", "List all saved profiles")
	fmt.Println("")

	fmt.Println("Action Commands:")
	fmt.Printf("  %-25s %s\n", "go <name>", "Launch a profile")
	fmt.Printf("  %-25s %s\n", "go -k <name>", "Kill current apps (Interactive), then launch")
	fmt.Printf("  %-25s %s\n", "go -fk <name>", "Force Kill current apps (Instant), then launch")
	fmt.Println("")
}

func runCalibrate() {
	fmt.Println("")
	fmt.Printf("Scanning system for noise...\n")

	pureNoiseList, err := sys.Capture()
	fmt.Println("")
	fmt.Printf("Caught %d ghost apps in the background\n", len(pureNoiseList))

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

	fmt.Println("")
	fmt.Printf("CALIBRATION DONE\n")
	fmt.Println("")
}

func runSave() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("Give the state a name...")
		fmt.Println("switchblade save <name>")

		return
	}

	profilname := os.Args[2]

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
	if err != nil {
		fmt.Printf("Error Saving the state: %v", err)
	}

	fmt.Printf("State Saved Successfully!\n")
	fmt.Println("-----------------------------")

	fmt.Printf("State Name: %s\n", profilname)
	for _, app := range clearList {
		app = filepath.Base(app)
		app = strings.TrimSuffix(app, ".exe")
		fmt.Printf("  - %s\n", app)
	}

	fmt.Println("")
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

func runList() {

	if len(os.Args) < 3 {
		PrintAllStates()
		return
	}

	stateName := os.Args[2]

	SavedProfileStruct, err := profile.LoadProfile(stateName)

	if err != nil {
		fmt.Printf("Error loading state to print: %v\n", err)
		return
	}

	fmt.Printf("%s:\n", SavedProfileStruct.Name)

	for _, file := range SavedProfileStruct.Apps {

		file = filepath.Base(file)
		file = strings.TrimSuffix(file, ".exe")

		fmt.Printf("%s\n", file)
	}

}

func PrintAllStates() {
	files, err := filepath.Glob("*.json")

	if err != nil {
		fmt.Printf(" Error listing files : %v", err)
	}

	if len(files) == 0 {
		fmt.Println("No profiles found. Save one with: switchblade save <name>")
		return
	}

	for _, file := range files {

		if file == "Noise.json" {
			continue
		}

		file = strings.TrimSuffix(file, ".json")

		fmt.Printf("%s\n", file)
	}
}
