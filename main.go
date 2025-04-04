package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func startServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server on port %d stopped: %v", port, err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <port1> <port2> ...", os.Args[0])
	}

	var wg sync.WaitGroup
	for _, arg := range os.Args[1:] {
		port, err := strconv.Atoi(arg)
		if err != nil || port < 1 || port > 65535 {
			log.Fatalf("Invalid port number: %s", arg)
		}
		wg.Add(1)
		go startServer(port, &wg)
	}
	wg.Wait()
}
