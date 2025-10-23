package main

import (
	"flag"
	"fmt"
	"main/net"
	"os"
	"strconv"
	"strings"
	"time"
)

var VERSION float32 = 0.1

func main() {

	dir, _ := os.UserConfigDir()

	// Parse command line args
	flag.Parse()
	option, param, file := flag.Arg(0), flag.Arg(1), flag.Arg(2)

	net.ReadWhitelist(dir + "/peer/whitelist.conf")

	// Send logic
	if option == "send" {
		dest := strings.Split(param, ":")
		addr, pseudo := dest[0], "unknown"
		port, _ := strconv.Atoi(dest[1])
		if !net.IsIP(addr) && net.PeerExists(addr) {
			pseudo = addr
			addr = net.GetIPByNickname(addr)
		}
		net.SendRequest(addr, strconv.Itoa(port), net.GetIPv4(), pseudo, net.GetFile(file))
		switch net.GetAnswer(strconv.Itoa(port+1), addr, strconv.Itoa(port)) {
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
		net.ReadWhitelist(dir + "/peer/whitelist.conf")
		fmt.Println(net.Whitelist)
	}

	// Help flag
	if option == "--help" || option == "-h" || option == "help" {
		fmt.Println("<-*->")
		fmt.Println("Send a file :")
		fmt.Println(" peer send <ip:port> <file>")
		fmt.Println(" -> peer send 10.41.230.165:9000 picture.png")
		fmt.Println("\nReceive a file :")
		fmt.Println(" peer receive <port>")
		fmt.Println(" -> peer receive 9000")
		fmt.Println("<-*->")
	}

	if option == "--version" || option == "-v" || option == "version" {
		fmt.Printf("\033[33m%v\033[0m\n", VERSION)
	}
}
