package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // Step 1: Create a new ServeMux

	server := &http.Server{ // Step 2: Create the built-in http.Server struct
		Addr:    ":8080", // - Set the address to ":8080"
		Handler: mux,     // - Use the mux as the handler
	}
	mux.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Printf("Server is listening on %s\n", server.Addr)
	err := server.ListenAndServe() // Step 3: Start the server
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
