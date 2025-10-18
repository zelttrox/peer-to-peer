package net

import (
	"fmt"
	"net"
)

// Connect to another peer using it's ip and open port
func Connect(ip string, port string) net.Conn {
	peer, err := net.Dial("tcp", ip+":"+port)
	if err != nil {fmt.Println(err)}
	fmt.Println("Connected to", ip)
	defer peer.Close()
	return peer
}

// Returns the local IPv4 address
func GetIPv4() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {fmt.Println(err)}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP.String()
}