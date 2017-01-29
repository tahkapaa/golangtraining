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
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	fmt.Printf("Connection from %v\n", conn.RemoteAddr())
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	var method string
	var uri string
	var body string
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			parts := strings.Split(ln, " ")
			if len(parts) > 1 {
				method = parts[0]
				uri = parts[1]
			} else {
				method = "Unknown"
				uri = "Unknown"
			}
		}
		if ln == "" {
			body = fmt.Sprintf("Method: %s, Uri: %s", method, uri)
			writeStatus(conn)
			writeHeader(conn, len(body))
			io.WriteString(conn, body)
			break
		}
		i++
	}
	fmt.Println(body)
	fmt.Println("Connection closed")
}

func writeStatus(w io.Writer) {
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
}

func writeHeader(w io.Writer, bodyLen int) {
	fmt.Fprintf(w, "Content-Length: %d\r\n", bodyLen)
	fmt.Fprint(w, "Content-Type: text/plain\r\n")
	fmt.Fprint(w, "\r\n")
}
