package net

import (
	"fmt"
	"net"
)

// Send a file to a peer
func SendFile(ip string, port string, file File) {

}

// Send a request to a peer
func SendRequest(ip string, port string, source string, file File) {
	peer, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
	}
	defer peer.Close()
	peer.Write([]byte(
		"Incoming file download received from " + source + "\n" +
			"File: " + file.Name + " -> " + file.Size + "\n" +
			"Download this file? (Y/n)",
	))
	fmt.Println("File download request sent. Waiting for answer from peer..")
}
