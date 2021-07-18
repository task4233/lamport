package main

import (
	"fmt"
	"net"
	"os"
)

// Server はIPアドレスとポート番号を保持します
type Server struct {
	ipv4 string
	port string
}

// Addr はIPアドレスとポート番号を結合したstringを返します
func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%s", s.ipv4, s.port)
}

func main() {
	srv := &Server{
		ipv4: "224.0.0.1",
		port: "56789",
	}

	fmt.Println("Receiver: ", srv.Addr())

	udpAddr, err := net.ResolveUDPAddr("udp", srv.Addr())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ResolveUDP: %s", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to ListenMulticastUDP: %s", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	buf := make([]byte, 2048)
	for {
		length, remoteAddr, err := listener.ReadFrom(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to ReadForm: %s", err.Error())
			continue
		}

		fmt.Println("Sender: ", remoteAddr.String())
		fmt.Println("Content: ", string(buf[:length]))

		go func() {
			conn, err := net.Dial("udp", srv.Addr())
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to Dial: %s", err.Error())
				os.Exit(1)
			}
			defer conn.Close()

			conn.Write(buf[:length])
			fmt.Println(string(buf[:length]))
		}()
	}

}
