package net

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var SourceIP string
var FileName string

// Open a specified port to receive files from peers
func OpenPort(port string) {
	listener, _ := net.Listen("tcp", ":"+port)
	fmt.Println("Opened port", port, "- waiting for requests from peers..")
	conn, _ := listener.Accept()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error", err)
	}
	PrintRequest(string(buffer[:n]))
}

// Print the request in the terminal
func PrintRequest(request string) {
	// file-name * file-size * source
	req := strings.Split(request, "*")
	fmt.Printf("Incoming file request from \x1b[36m%s\x1b[0m\n", req[2])
	fmt.Printf("\x1b[33m%s -> %s\x1b[0m\n", req[0], req[1])
	FileName = req[0]
	SourceIP = req[2]
}

// Accept or deny file transfer
func SendAnswer(source string, port string) {
	var input string
	for {
		fmt.Print("Accept file transfer? (Y/n): ")
		fmt.Scanln(&input)
		ans := strings.ToLower(strings.TrimSpace(input))
		if ans != "y" && ans != "n" {
			continue
		}
		conn, err := net.Dial("tcp", source+":"+port)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer conn.Close()
		conn.Write([]byte(ans))
		break
	}
	if input == "n" {
		os.Exit(0)
	}
}

// Wait for peer answer after sending a file download request
func GetAnswer(port string) bool {
	listener, _ := net.Listen("tcp", ":"+port)
	fmt.Println("File transfer request sent - waiting for answer from peer..")
	conn, _ := listener.Accept()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
	} else if string(buffer[:n]) == "y" {
		fmt.Println("File transfer accepted by peer, proceeding..")
		return true
	} else if string(buffer[:n]) == "n" {
		fmt.Println("File transfer denied by peer, cancelling..")
		return false
	} else {
		fmt.Println("Unkown answer from peer, cancelling..")
		return false
	}
	return false
}

// Create a file and receive it's content
func ReceiveFile(port string, file string) {
	output, err := os.Create(file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer output.Close()
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error:", err)
	}
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	io.Copy(output, conn)
}
