package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	ipv4Addr := "224.0.0.1"
	port := ":56789"
	addr := ipv4Addr + port
	waitTime := 1
	format := "Message "

	fmt.Println("Sender: ", addr)
	conn, err := net.Dial("udp", addr)
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
