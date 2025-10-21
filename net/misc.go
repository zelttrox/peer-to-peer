package net

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
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

func ProgressBar(total int, progress int) string {
	return ""
}

func ByteSize(size string) int32 {
	sizeStr, _ := strings.CutSuffix(size, " Mo")
	floatSize, _ := strconv.ParseFloat(sizeStr, 64)
	return int32(floatSize * 1024 * 1024)
}

func GetIcon(ext string) string {
	decomp := strings.Split(ext, ".")
	ext = decomp[1]
	switch ext {
	case "png":
		return ""
	case "pdf":
		return ""
	case "iso":
		return ""
	case "exe":
		return ""
	case "go":
		return ""
	case "js":
		return ""
	case "java":
		return ""
	case "ts":
		return ""
	case "conf":
		return ""
	case "py":
		return ""
	case "rs":
		return ""
	case "mp4":
		return ""
	default:
		return ""
	}
}