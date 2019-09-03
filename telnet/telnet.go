package telnet

import (
	"log"
	"os"

	"github.com/ircop/tclient"
)

// MyTelnet struct
type MyTelnet struct {
	client *tclient.TelnetClient
}

// New Connect to model telnet api
func New() (*MyTelnet, error) {
	log.Println("Trying to connect...")

	telnet := tclient.New(3, os.Getenv("PASSWORD"), "", "")
	telnet.SetLoginPrompt(`[Pp]ass[Ww]ord\:$`)
	err := telnet.Open(os.Getenv("IP"), 23)

	if err != nil {
		return nil, err
	}

	log.Println("Connected")

	return &MyTelnet{telnet}, nil
}

// ChangeIP ...
func (m *MyTelnet) ChangeIP() {
	m.client.Cmd(`wan node index 3`)
	m.client.Cmd(`wan node disable`)
	m.client.Cmd(`wan node enable`)
	m.client.Cmd(`wan node save`)
}
