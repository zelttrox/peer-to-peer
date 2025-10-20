package main

import (
	"flag"
	"fmt"
	"main/config"
	"main/net"
	"strings"
)

func main() {

	// Parse command line args
	flag.Parse()
	option, param, file := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	config.ReadWhitelist("config/whitelist.conf")

	// Send logic
	if option == "send" {
		dest := strings.Split(param, ":")
		addr, port, pseudo := dest[0], dest[1], "unknown"
		if !net.IsIP(addr) && config.PeerExists(addr) {
			pseudo = addr
			addr = config.GetIPByNickname(addr)
		}
		fmt.Println("Addr:", addr)
		net.SendRequest(addr, port, net.GetIPv4(), pseudo, net.GetFile(file))
		switch net.GetAnswer("9334") {
		case true:
			net.SendFile(addr, "9339", net.GetFile(file).Path)
		case false:
			return
		}
	}

	// Receive logic
	if option == "receive" {
		net.OpenPort(param)
		net.SendAnswer(net.SourceIP, "9334")
		net.ReceiveFile("9339", net.FileName)
	}

	// Whitelist logic
	// whitelist add|remove|edit|list 192.168.10.27 nickname
	if option == "whitelist" {
		switch param {
		case "list":
			config.ReadWhitelist("config/whitelist.conf")
			fmt.Println(config.Whitelist)
		case "add":
			config.AddPeer(config.Peer{Nickname: flag.Arg(2), IP: flag.Arg(3)})
			fmt.Println(config.Whitelist)
		}
	}

	// Help flag
	if option == "--help" || option == "-h" || option == "help" {
		fmt.Println("<-*->")
		fmt.Println("Send a file :")
		fmt.Println(" peer send <ip:port> <file>")
		fmt.Println(" -> peer send 10.41.230.165:9000 \"picture.png\"")
		fmt.Println("\nReceive a file :")
		fmt.Println(" peer receive <port>")
		fmt.Println(" -> peer receive 9000")
		fmt.Println("<-*->")
	}
}
