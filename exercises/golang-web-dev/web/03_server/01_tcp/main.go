package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("Connection from %v\n", conn.RemoteAddr())
	defer conn.Close()
	io.WriteString(conn, "I see you connected\n")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "BYE" {
			break
		}
		fmt.Fprintf(conn, "Why '%s'\n", ln)
	}
	io.WriteString(conn, "Bye bye!\n")
	fmt.Println("Connection closed")
}
