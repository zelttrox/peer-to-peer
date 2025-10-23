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

	buffer := make([]byte, 32*1024) // 32KB chunks
	var sent int64
	total := ByteSize(GetFile(path).Size)

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			written, _ := conn.Write(buffer[:n])
			sent += int64(written)
			fmt.Printf("\r%.2f%%", float64(sent)*100/float64(total))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	fmt.Printf("\r")
	fmt.Println("\x1b[32mFile transfer complete!\x1b[0m")

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
