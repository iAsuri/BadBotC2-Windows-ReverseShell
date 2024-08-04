package commands

// Command Structure

type command struct {
	Name, Description string
	MaxArgs           int
	Action            func(...string)
}

//  Command Structure Array
var Commands []command
