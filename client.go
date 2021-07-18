package lamport

import (
	"errors"
	"fmt"
	"net"
)

const (
	BASE_IPV4 = "224.0.0.1"
	BASE_PORT = 56789
)

var isUsedErr = errors.New("this port is used")

type Client struct {
	id   int // 0-indexed
	port int
	tick int
}

func (c *Client) Addr() string {
	return fmt.Sprintf("%s:%d", BASE_IPV4, BASE_PORT)
}

func (c *Client) checkUsed() error {
	conn, err := net.Dial("udp", c.Addr())
	if err != nil {
		return fmt.Errorf("failed to Dial: %w", err)
	}
	defer conn.Close()

	return nil
}

// NewClient checks other clients
// with multicast REQ message
func NewClient() *Client {
	client := &Client{
		id: 0,

		port: BASE_PORT,
		tick: 0,
	}

	c.clients = append(c.clients, client)
	return client
}

func (c *Client) Run() int {

	return 0
}
