package main

import (
	"fmt"
	"net/http"
)


func main() {
	// Start an HTTP server
	http.HandleFunc("/smarg/audience", audienceHandler)

	fmt.Println("Server is listening on port 3002...")
	err := http.ListenAndServe(":3002", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}