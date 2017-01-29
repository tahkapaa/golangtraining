package main

import (
	"bufio"
	"fmt"
	"io"
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
	i := 0
	var resp string
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			words := strings.Split(line, " ")
			resp = getResponse(words[0])
		}
		if line == "" {
			break
		}
		fmt.Println(line)
		i++
	}
	writeResponse(conn, resp)
}

func writeResponse(w io.Writer, body string) {
	fmt.Println("BODY:", body)
	io.WriteString(w, "HTTP/1.1 200 OK\r\n")            // status line
	fmt.Fprintf(w, "Content-Length: %d\r\n", len(body)) // header
	fmt.Fprint(w, "Content-Type: text/plain\r\n")       // header
	io.WriteString(w, "\r\n")
	io.WriteString(w, body)
}

func getResponse(method string) string {
	switch method {
	case "GET":
		return "Here you go!"
	case "POST":
		return "Thank you so much!"
	}
	return "And you are?"
}
