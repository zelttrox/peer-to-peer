package main

import (
	"flag"
	"fmt"
	"main/net"
	"strings"
)

func main() {

	// Parse command line args
	flag.Parse()
	option, param, file := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	// Send logic
	if (option == "send") {
		dest := strings.Split(param, ":")
		ip, port := dest[0], dest[1]
		peer := net.Connect(ip, port)
		fmt.Println("Source IP", net.GetIPv4())
		net.SendRequest(peer, net.GetIPv4(), net.GetFile(file))
	}

	// Receive logic
	if (option == "receive") {
		// Server logic
		fmt.Println(file)
	}
}