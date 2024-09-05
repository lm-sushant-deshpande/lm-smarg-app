package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const validAPIKey = "SMARG6002A2024LEMMA" // Replace with your actual API key
const filePath = "/etc/appdmin/app/smarg/data"

func audienceHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method, only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Authenticate API_Key from the request header
	apiKey := r.Header.Get("API_Key")
	if apiKey == "" {
		http.Error(w, "API_Key is missing", http.StatusUnauthorized)
		fmt.Println("Unauthorized: API_Key is missing")
		return
	}

	if apiKey != validAPIKey {
		http.Error(w, "Invalid API_Key", http.StatusUnauthorized)
		fmt.Println("Unauthorized: Invalid API_Key")
		return
	}

	// Generate a timestamped filename
	timestamp := time.Now().Format("20060102_150405") // Format as YYYYMMDD_HHMMSS
	fileName := fmt.Sprintf("request_body_%s.txt", timestamp)

	// Create the file with the timestamped name
	file, err := os.Create(filePath + fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the request body directly to the file
	_, err = io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		fmt.Println("Error writing to file:", err)
		return
	}
	defer r.Body.Close()

	// Send a success message back to the API caller
	fmt.Fprintf(w, "Request body successfully written to %s\n", fileName)
	fmt.Printf("Request body successfully written to %s\n", fileName)
}
