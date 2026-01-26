package main

import (
	"github.com/AY88o/switchblade/internal/profile"
)

func main() {

	school := profile.Profile{
		Name: "STUDY MODE",
		Apps: []string{"notepad", "calc"},
	}

	school.Start()
}
