package slave

import "net"

// Client Structure
type Client struct {
	Name    string
	Session net.Conn

	// Xor Key
	Xorkey int
}
