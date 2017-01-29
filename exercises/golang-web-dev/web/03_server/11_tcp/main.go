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
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	fmt.Printf("Connection from %v\n", conn.RemoteAddr())
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			body := "CHECK THE PAYLOAD"
			writeStatus(conn)
			writeHeader(conn, len(body))
			io.WriteString(conn, body)
			break
		}
	}
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
