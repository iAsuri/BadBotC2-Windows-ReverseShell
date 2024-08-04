package commands

import (
	"log"
	"strings"
)

// Command register Used to initiate Commands for Admins!
func Init_Commands() {

	Commands = append(Commands, command{Name: "help", Description: "This Command !", Action: help, MaxArgs: 1})
	Commands = append(Commands, command{Name: "ls", Description: "Shows Infected Users", Action: shellCount, MaxArgs: 1})
	Commands = append(Commands, command{Name: "cd", Description: "Opens Clients Enviorment | syntax cd ClientNameHere", Action: useShell, MaxArgs: 2})
	Commands = append(Commands, command{Name: "cls", Description: "Clears Screen", Action: ClearScreen, MaxArgs: 0})

	log.Println("[BadBot] Loaded Commands > ", len(Commands))

}

func Findcommand(input string) bool {

	input = strings.ToLower(input)

	Args := strings.Split(input, " ")

	for i := 0; i < len(Commands); i++ {

		if Commands[i].Name == Args[0] && len(Args) == Commands[i].MaxArgs {
			Commands[i].Action(Args...)
			return true
		}

	}

	// No Commands Found return error
	return false

}
