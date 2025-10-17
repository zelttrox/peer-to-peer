package main

import (
	"flag"
	"fmt"
)

func main() {

	// Parse command line args
	flag.Parse()

	// Handle args
	if (flag.Arg(0) == "send") {
		fmt.Println("send")
		// Client logic
	}
	if (flag.Arg(0) == "receive") {
		fmt.Println("receive")
		// Server logic
	}
}