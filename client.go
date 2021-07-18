package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	REQ_FORMAT = "REQ"
	ACK_FORMAT = "ACK"
	CLIENT_NUM = 2
)

// Client はIPアドレスとポート番号を保持します
type Client struct {
	ipv4 string
	port string
	id   int
	tick int
}

// Addr はIPアドレスとポート番号を結合したstringを返します
func (c *Client) Addr() string {
	return fmt.Sprintf("%s:%s", c.ipv4, c.port)
}

func (c *Client) SendREQ() error {
	conn, err := net.Dial("udp", c.Addr())
	if err != nil {
		return fmt.Errorf("failed to Dial: %w", err)
	}
	defer conn.Close()

	mes := fmt.Sprintf("%s%d:%d.%d", REQ_FORMAT, c.id, c.tick, c.id)
	conn.Write([]byte(mes))
	c.tick++

	fmt.Printf("[ID%d](t=%d) %s\n", c.id, c.tick, mes)

	return nil
}

func (c *Client) SendACK(recvMes []byte) error {
	targetID, err := strconv.Atoi(string(recvMes[:bytes.IndexRune(recvMes, ':')]))
	if err != nil {
		return fmt.Errorf("failed to strconv.Atoi: %w", err)
	}

	conn, err := net.Dial("udp", c.Addr())
	if err != nil {
		return fmt.Errorf("failed to Dial: %w", err)
	}
	defer conn.Close()

	// ACKtarget-from:tick.id
	mes := fmt.Sprintf("%s%d-%d:%d.%d", ACK_FORMAT, targetID, c.id, c.tick, c.id)
	conn.Write([]byte(mes))
	c.tick++

	fmt.Printf("[ID%d](t=%d) %s\n", c.id, c.tick, mes)

	return nil
}

func (c *Client) ExecuteTask() {
	fmt.Printf("[ID%d](t=%d) Execute Task -> %d\n", c.id, c.tick, c.id)
	c.tick++
}

func (c *Client) AdjustTick(receivedTick int) {
	if c.tick <= receivedTick {
		c.tick = receivedTick + 1
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, len(os.Args))
		os.Exit(1)
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "./client [id]")
		os.Exit(1)
	}

	srv := &Client{
		ipv4: "224.0.0.1",
		port: "56789",
		id:   id,
		tick: 0,
	}

	fmt.Printf("[ID%d] Begin on %s\n", srv.id, time.Now().Format("15:04:05"))

	checked := 0

	udpAddr, err := net.ResolveUDPAddr("udp", srv.Addr())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ResolveUDP: %s\n", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ListenMulticastUDP\n: %s", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	go func() {
		for {
			buf := make([]byte, 2048)
			_, _, err := listener.ReadFrom(buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to ReadForm: %s\n", err.Error())
				break
			}

			tick, err := strconv.Atoi(string(buf[bytes.IndexRune(buf, ':')+1 : bytes.IndexRune(buf, '.')]))
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed strconv.Atoi: %s\n", err.Error())
				continue
			}
			srv.AdjustTick(tick)

			// fmt.Printf("[CLIENT] Tick: %d\n", srv.tick)
			// fmt.Println("Checked: ", checked)
			// fmt.Println("Sender: ", remoteAddr.String())
			// fmt.Println("Content: ", string(buf[:length]))

			if isACK(buf[:3]) {
				toID, err := strconv.Atoi(string(buf[3:bytes.IndexRune(buf, '-')]))
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed strconv.Atoi: %s\n", err.Error())
					continue
				}

				if toID != srv.id {
					continue
				}

				fromID, err := strconv.Atoi(string(buf[bytes.IndexRune(buf, '-')+1 : bytes.IndexRune(buf, ':')]))
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed strconv.Atoi: %s\n", err.Error())
					continue
				}

				checked |= 1 << fromID
				if checked == (1<<CLIENT_NUM)-1 {
					srv.ExecuteTask()
					checked = 0
				}
				continue
			}

			if isREQ(buf[:3]) {
				if err := srv.SendACK(buf[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "failed to SendACK: %s\n", err.Error())
				}
				continue
			}

			if err := srv.SendREQ(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to SendREQ: %s\n", err.Error())
				break
			}
		}
	}()

	for idx := 0; idx < 3; idx++ {
		time.Sleep(3 * time.Second)
		// fmt.Println("Sender: ", srv.Addr())
		if err := srv.SendREQ(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to SendREQ: %s\n", err.Error())
			os.Exit(1)
		}

	}

	fmt.Printf("[ID%d] End on %s\n", srv.id, time.Now().Format("15:04:05"))

}
