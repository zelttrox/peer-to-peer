package net

import (
	"fmt"
	"net"
	"regexp"
)

// Returns the local IPv4 address
func GetIPv4() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP.String()
}

// Check if a string is an IP
func IsIP(expr string) bool {
	var ip = `^[0-9].[0-9].[0-9].[0-9]$`
	isIP, _ := regexp.MatchString(ip, expr)
	if isIP {
		return true
	} else {
		return false
	}
}
