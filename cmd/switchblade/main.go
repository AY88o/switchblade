package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println("Usage:")
		fmt.Println("Switchblade Calibrate  ....(calibrate the tool)")
		fmt.Println("OR")
		fmt.Println("Switchblade Save <name>   ...(Save state)")

		return

	}

}
