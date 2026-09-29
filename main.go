package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Dial failed: %v", err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET /bar HTTP/1.0\r\n\r\n")

	buf := bufio.NewReader(conn)
	status, err := buf.ReadString('\n')
	if err != nil {
		log.Fatal("readString error: %v", err)
	}
	fmt.Printf("status: %s", status)

	// Header parsing
	headers := make(map[string]string)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			log.Fatalf("header read error: %v", err)
		}

		trimmed := strings.TrimRight(line, "\r\n")

		if trimmed == "" {
			break
		}

		parts := strings.SplitN(trimmed, ": ", 2)
		if len(parts) == 2 {
			headers[parts[0]] = parts[1]
		}
	}

	fmt.Println("headers:")
	for k, v := range headers {
		fmt.Printf("\n%s: %s", k, v)
	}

	// body parsing
	body, err := io.ReadAll(buf)
	if err != nil {
		log.Fatalf("body read error: %s", err)
	}
	fmt.Printf("\nbody: %s", string(body))
}
