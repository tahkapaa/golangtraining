package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var host string
	for scanner.Scan() {
		ln := scanner.Text()
		if strings.Contains(strings.ToLower(ln), "host:") {
			host = strings.Fields(ln)[1]
			break
		}
	}
	conn.Write([]byte(host))
}
