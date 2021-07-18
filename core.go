package lamport

import (
	"fmt"
	"net"
)

// Core mangages clients
type Core struct {
	clients []*Client
}

func NewCore(ipv4 string) *Core {
	return &Core{
		clients: []*Client{},
	}
}

func (c *Core) Run() error {
	listener, err := net.Listen("tcp", c.Addr())
	if err != nil {
		return fmt.Errorf("failed net.Dial: %w", err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return fmt.Errorf("failed Accept: %w", err)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 128)

		if _, err := conn.Read(buf); err != nil {
			return
		}

		fmt.Printf("buf: %s\n", string(buf))
	}()

	return nil
}
