package admin

import (
	"BadBot/source/admin/commands"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func AdminHandle() {
	var Input string

	time.Sleep(30 * time.Millisecond)

	log.Println("[BadBot] Attaching to Stdin...")

	time.Sleep(30 * time.Millisecond)

	// End of Terminal Art =--------------------------------------------

	// Create a new scanner to read from Stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Infinite loop to continuously prompt the user
	for {
		// Print the command prompt

		fmt.Print("BadBot (EvilEye/RevShell) $ ")

		// Read user input using the scanner
		if !scanner.Scan() {
			break
		}

		Input = scanner.Text()

		Input = strings.TrimSpace(Input)

		if len(Input) == 0 {
			continue
		}

		// Check if the command exists using the commands package
		if !commands.Findcommand(Input) {
			fmt.Printf("'%s' is not recognized as an internal command\n", Input)
		}

	}
}
