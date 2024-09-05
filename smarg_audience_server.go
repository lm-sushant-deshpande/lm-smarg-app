package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func audienceHandler(w http.ResponseWriter, r *http.Request) {
	// URL to make the request
	url := "https://smargtechnologies.in:3015/smarg_ai/siteConfiguration/125"

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Error creating HTTP request", http.StatusInternalServerError)
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Add the API_Key header
	req.Header.Set("API_Key", "SMARG6002A2024LEMMA") // Replace with your actual API key

	// Send the request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error making HTTP request", http.StatusInternalServerError)
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error: received non-200 response code %d", resp.StatusCode), http.StatusInternalServerError)
		fmt.Printf("Error: received non-200 response code %d\n", resp.StatusCode)
		return
	}

	// Read the response body and write it to a file
	fileName := "response.txt"
	file, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the response body directly to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		fmt.Println("Error writing to file:", err)
		return
	}

	// Send a success message back to the API caller
	fmt.Fprintf(w, "Response successfully written to %s\n", fileName)
	fmt.Printf("Response successfully written to %s\n", fileName)
}
