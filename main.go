package main

import (
	"FlickUp/server"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}

	errs := make(chan error, 2)
	server := server.NewServer()

	go func() {
		fmt.Printf("Listening on port %s\n", addr)
		errs <- http.ListenAndServe(":"+addr, server)
	}()

	if err := <-errs; err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
