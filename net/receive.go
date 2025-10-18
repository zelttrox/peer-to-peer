package net

import (
	"fmt"
	"net"
)

// Open a specified port to receive files from peers
func OpenPort(port string) {
	listener, _ := net.Listen("tcp", ":"+port)
	fmt.Println("Opened port", port, "- waiting for requests from peers..")
	conn, _ := listener.Accept()
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {fmt.Println("Error", err)}
	fmt.Println("Received:", string(buf[:n]))
}
