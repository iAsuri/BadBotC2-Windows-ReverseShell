package commands

import "fmt"

func help(args ...string) {

	if len(args) != 1 {
		return
	}

	fmt.Printf("\t----- Help -----\n       Loaded: %d Commands\n\n", len(Commands))

	// Iterate Over Commands Array Print and print to screen
	for i := 0; i < len(Commands); i++ {
		fmt.Printf("%-15s | %s\n", Commands[i].Name, Commands[i].Description)
	}

	fmt.Print("\n")
}
