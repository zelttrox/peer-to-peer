package main

import (
	"flag"
	"fmt"
	"main/config"
	"main/net"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Parse command line args
	flag.Parse()
	option, param, file := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	config.ReadWhitelist("config/whitelist.conf")

	// Send logic
	if option == "send" {
		dest := strings.Split(param, ":")
		addr, pseudo := dest[0], "unknown"
		port, _ := strconv.Atoi(dest[1])
		if !net.IsIP(addr) && config.PeerExists(addr) {
			pseudo = addr
			addr = config.GetIPByNickname(addr)
		}
		fmt.Println(net.Byte(net.GetFile(file).Size))
		net.SendRequest(addr, strconv.Itoa(port), net.GetIPv4(), pseudo, net.GetFile(file))
		switch net.GetAnswer(strconv.Itoa(port + 1)) {
		case true:
			time.Sleep(300 * time.Millisecond)
			net.SendFile(addr, strconv.Itoa(port+2), net.GetFile(file).Path)
		case false:
			return
		}
	}

	// Receive logic
	if option == "receive" {
		port, _ := strconv.Atoi(param)
		net.OpenPort(param)
		net.SendAnswer(net.SourceIP, strconv.Itoa(port+1))
		net.ReceiveFile(strconv.Itoa(port+2), net.FileName)
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

	if option == "prog" {
		net.ProgressBar(0, 0)
	}
}
