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
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	fmt.Printf("Waiting for requests on \x1b[36m%s\x1b[0m:\x1b[36m%s\x1b[0m\n", GetIPv4(), port)
	conn, _ := listener.Accept()
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error", err)
	}
	request := strings.Split(string(buffer[:n]), "*")
	if PeerBlocked(request[2]) {
		return
	} else {
		PrintRequest(request)
	}
	
}

// Print the request in the terminal
func PrintRequest(req []string) {
	// file-name * file-size * source
	fmt.Printf("Incoming file request from \x1b[36m%s\x1b[0m (%s) \n", req[2], req[3])
	fmt.Printf("󱞩 \x1b[94m%s %s \x1b[0m-> \x1b[94m%s\x1b[0m\n", GetIcon(req[0]), req[0], req[1])
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
func GetAnswer(port string, destIP string, destPort string) bool {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	fmt.Printf("File request sent to \x1b[36m%s\x1b[0m:\x1b[36m%s\x1b[0m.\n", destIP, destPort)
	fmt.Println("Waiting for answer from peer..")
	conn, _ := listener.Accept()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
	} else if string(buffer[:n]) == "y" {
		return true
	} else if string(buffer[:n]) == "n" {
		fmt.Println("\x1b[31mFile transfer denied by peer.\x1b[0m")
		return false
	} else {
		fmt.Println("\x1b[31mUnexpected answer '", string(buffer[:n]), "' from peer\x1b[0m")
		return false
	}
	defer conn.Close()
	return false
}

// Create a file and receive it's content
func ReceiveFile(port string, file string) {
	output, err := os.Create(file)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer output.Close()
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer listener.Close()
	
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	io.Copy(output, conn)
}
