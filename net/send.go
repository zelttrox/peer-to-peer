package net

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Send a file to a peer
func SendFile(ip string, port string, path string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	io.Copy(conn, file)
}

// Send a request to a peer
func SendRequest(ip string, port string, source string, file File) {
	peer, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
	}
	peer.Write([]byte(file.Name + "*" + file.Size + "*" + source))
}
