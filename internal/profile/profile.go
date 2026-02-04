package profile

import (
	"fmt"
	"os/exec"
	"time"
)

type Profile struct {
	Name string
	Apps []string
}

func (p Profile) Start() {
	fmt.Printf("\n--- IGNITING %s PROTOCOL ---\n", p.Name)

	for _, app := range p.Apps {
		fmt.Printf("[+] Launching %s...\n", app)
		cmd := exec.Command("cmd", "/C", "start", "", app)
		cmd.Start()
		time.Sleep(500 * time.Millisecond)
	}
}
