package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	// Start an HTTP server
	http.HandleFunc("/smarg/audience", audienceHandler)

	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}