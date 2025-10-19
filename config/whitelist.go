package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Peer struct {
	Nickname string
	IP       string
}

var Whitelist []Peer

// Read the whitelist
func ReadWhitelist(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		Whitelist = append(Whitelist, Peer{Nickname: parts[0], IP: parts[1]})
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// Check if a peer exists
func PeerExists(nickname string) bool {
	for _, peer := range Whitelist {
		if peer.Nickname == nickname {
			return true
		} else {
			continue
		}
	}
	return false
}

// Return the nickname of a peer based on a given IP
func GetNicknameByIP(ip string) string {
	var nickname string
	for _, peer := range Whitelist {
		if peer.IP == ip {
			nickname = peer.Nickname
		}
	}
	return nickname
}

// Return the IP of a peer based on a given nickname
func GetIPByNickname(nickname string) string {
	var ip string
	for _, peer := range Whitelist {
		if peer.Nickname == nickname {
			ip = peer.IP
		}
	}
	return ip
}

// Add a peer to whitelist
func AddPeer(peer Peer) {
	Whitelist = append(Whitelist, peer)
}