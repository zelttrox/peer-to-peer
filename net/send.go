package net

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

// Turn a filepath into a File type variable
func GetFile(path string) File {
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	return File{stat.Name(), path, fmt.Sprint(stat.Size()), filepath.Ext(path)}
}

// Send a file to a peer
func SendFile(ip string, port string, file File) {

}

// Send a request to a peer
func SendRequest(peer net.Conn, source string, file File) {
	peer.Write([]byte(
		"Incoming file download received from " + source + "\n" +
		"File: " + file.Name + "." + file.Type + " " + file.Type + "\n" +
		"Download this file? (Y/n)",
	))
	fmt.Println("File download request sent. Waiting for answer from peer..")
}
