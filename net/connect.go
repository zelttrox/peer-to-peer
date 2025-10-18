package net

import (
	"fmt"
	"net"
)

// Returns the local IPv4 address
func GetIPv4() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {fmt.Println(err)}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP.String()
}