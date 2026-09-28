package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/bar")
	if err != nil {
		log.Fatal("http Get failure: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll failure: %s", err)
	}
	fmt.Printf("body: %s\n", body)
}
