package main

import (
	"flag"
	"main/net"
	"strings"
)

func main() {

	// Parse command line args
	flag.Parse()
	option, param, file := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	// Send logic
	if option == "send" {
		dest := strings.Split(param, ":")
		ip, port := dest[0], dest[1]
		net.SendRequest(ip, port, net.GetIPv4(), net.GetFile(file))
		switch net.GetAnswer("9300") {
		case true:
			net.SendFile(ip, port, net.GetFile(file))
		case false:
			return
		}
	}

	// Receive logic
	if option == "receive" {
		net.OpenPort(param)
		net.SendAnswer(net.SourceIP, "9300")
	}
}
