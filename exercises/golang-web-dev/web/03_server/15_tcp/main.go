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
	var method, uri string
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			parts := strings.Fields(ln)
			if len(parts) > 1 {
				method = parts[0]
				uri = parts[1]
			}
		}
		if ln == "" {

			break
		}
		i++
	}
	mux(conn, method, uri)
	fmt.Println("Connection closed")
}

func mux(w io.Writer, method string, uri string) {
	fmt.Println("Method, URI", method, uri)
	switch uri {
	case "/":
		index(w, method)
	case "/apply":
		apply(w, method)
	default:
		notFound(w)
		writeHeader(w, 0)
	}

}

func apply(w io.Writer, method string) {
	switch method {
	case "GET":
		getApply(w)
	case "POST":
		postApply(w)
	default:
		invalidMethod(w)
	}
}

func index(w io.Writer, method string) {
	switch method {
	case "GET":
		getIndex(w)
	default:
		invalidMethod(w)
	}
}

func getIndex(w io.Writer) {
	body := fmt.Sprint("<!DOCTYPE html><html><head></head><body><strong>HOLY COW THIS IS LOW LEVEL</strong></body></html>")
	writeResponse(w, body)
}

func getApply(w io.Writer) {
	body := fmt.Sprint("<!DOCTYPE html><html><head></head><body><strong>GET APPLY</strong></body></html>")
	writeResponse(w, body)
}

func postApply(w io.Writer) {
	body := fmt.Sprint("<!DOCTYPE html><html><head></head><body><strong>POST APPLY</strong></body></html>")
	writeResponse(w, body)
}

func writeResponse(w io.Writer, body string) {
	writeStatus(w)
	writeHeader(w, len(body))
	io.WriteString(w, body)
}

func writeStatus(w io.Writer) {
	fmt.Fprint(w, "HTTP/1.1 200 OK\r\n")
}

func notFound(w io.Writer) {
	fmt.Fprint(w, "HTTP/1.1 404 NOT FOUND\r\n")
	body := "404 NOT FOUND"
	writeHeader(w, len(body))
	writeResponse(w, body)
}

func invalidMethod(w io.Writer) {
	fmt.Fprint(w, "HTTP/1.1 405 INVALID METHOD\r\n")
	body := "405 INVALID METHOD"
	writeHeader(w, len(body))
	writeResponse(w, body)
}

func writeHeader(w io.Writer, bodyLen int) {
	fmt.Fprintf(w, "Content-Length: %d\r\n", bodyLen)
	fmt.Fprint(w, "Content-Type: text/html\r\n")
	fmt.Fprint(w, "\r\n")
}
