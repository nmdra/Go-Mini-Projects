package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	socketPath := "/tmp/unix_socket_http_example.sock"

	// Remove the socket file if it already exists
	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	// Create a Unix domain socket listener
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return
	}
	defer listener.Close()

	// HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request:")
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Unix socket HTTP server!"}`))
	})

	fmt.Println("Unix socket HTTP server is listening...")
	// Start the HTTP server using the Unix socket listener
	if err := http.Serve(listener, nil); err != nil {
		fmt.Printf("Error starting HTTP server: %v\n", err)
	}
}
