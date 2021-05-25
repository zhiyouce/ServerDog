package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {

	var hostAddress = flag.String("ip", "127.0.0.1", "target address")
	var port = flag.Int("port", 22, "port to run ssh")
	var user = flag.String("user", "ubuntu", "user for ssh login")
	var command = flag.String("command", "echo 'hello'", "command to run")
	flag.Parse()
	fmt.Println(*hostAddress)
	if net.ParseIP(*hostAddress) == nil {
		log.Fatalln("this is a valid ip address!")
	}
	fmt.Println("Start to code and have a nice day!")
}
