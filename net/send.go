package net

import (
	"fmt"
	"io"
	"net"
	"os"
)

var Total float64
var Progress float64

// Send a file to a peer
func SendFile(ip string, port string, path string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	io.Copy(conn, file)
}

// Send a request to a peer
func SendRequest(ip string, port string, source string, pseudo string, file File) {
	peer, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer peer.Close()

	peer.Write([]byte(file.Name + "*" + file.Size + "*" + source + "*" + pseudo))
}
