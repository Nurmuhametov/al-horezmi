package gaming

import (
	"fmt"
	"net"
)

type Communicator struct {
	active                bool
	conn                  net.Conn
	DataReceivedListeners []func(string2 string)
}

func (c Communicator) Start() {
	c.active = true
	go c.communicate()
}

func (c Communicator) communicate() {
	for {
		if !c.active { break }
		buff := make([]byte, 1024)
		_, err := c.conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading buffer")
			c.active = false
			break
		}
		data := string(buff)

		for _, value := range c.DataReceivedListeners {
			value(data)
		}
	}
}
func (c Communicator) Stop() {
	c.active = false
}
