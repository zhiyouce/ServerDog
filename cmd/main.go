package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type config struct {
	user    string
	address string
	command string
	port    int // int or string?
}

func main() {

	// Set the cli params
	var hostAddress = flag.String("ip", "127.0.0.1", "target address")
	// var port = flag.Int("port", 22, "port to run ssh")
	// var user = flag.String("user", "ubuntu", "user for ssh login")
	// var command = flag.String("command", "echo 'hello'", "command to run")
	flag.Parse()
	fmt.Println(*hostAddress)
	if net.ParseIP(*hostAddress) == nil {
		log.Fatalln("this is a valid ip address!")
	}

	// Set the env
	fmt.Printf("uid : %d\n", os.Getuid())
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(dir)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(pwd)

	// Downgrade this process's privilege
	// after reading the config
	fmt.Println("Start to code and have a nice day!")
}
