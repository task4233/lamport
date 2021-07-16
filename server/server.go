package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ipv4Addr := "224.0.0.1"
	port := ":56789"
	addr := ipv4Addr + port
	fmt.Println("Receiver: ", addr)

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
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
	}

}
