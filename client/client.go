package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// Client はIPアドレスとポート番号を保持します
type Client struct {
	ipv4 string
	port string
	id   int
	tick int
}

func (s *Client) Run() error {

	return nil
}

// Addr はIPアドレスとポート番号を結合したstringを返します
func (s *Client) Addr() string {
	return fmt.Sprintf("%s:%s", s.ipv4, s.port)
}

const (
	waitTime = 1
	format   = "Message "
)

func main() {
	client := &Client{
		ipv4: "224.0.0.1",
		port: "56789",
	}
	fmt.Println("Sender: ", client.Addr())

	conn, err := net.Dial("udp", client.Addr())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to Dial: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	c := 0
	for {
		time.Sleep(time.Duration(waitTime) * time.Second)
		message := format + strconv.Itoa(c)
		conn.Write([]byte(message))
		fmt.Println(message)

		c++
	}
}
