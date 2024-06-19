package main

import (
    "encoding/json"  // Importing encoding/json package to handle JSON encoding and decoding
    "net/http"       // Importing net/http package for HTTP server functionalities
    "log"            // Importing log package for logging
)

// Response struct defines the structure of the JSON response
type Response struct {
    Message string `json:"message"`  // Define a field `Message` in JSON with key "message"
}

// handler function handles the HTTP request to `/api/message` endpoint
func handler(w http.ResponseWriter, r *http.Request) {
    // Set response headers to indicate JSON content type and allow CORS
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // Allow GET and OPTIONS methods
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // Allow Content-Type header

    // Check if the request method is OPTIONS (preflight request for CORS)
    if r.Method == http.MethodOptions {
        return // If it's OPTIONS, return without processing further
    }

    // Create a Response object with a sample message
    response := Response{Message: "Hello from Go backend!"}

    // Encode the Response object as JSON and write it to the response writer
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Register handler function to handle requests to `/api/message` endpoint
    http.HandleFunc("/api/message", handler)

    // Log a message indicating the server is starting on port 8080
    log.Println("Starting server on :8080")

    // Start the HTTP server on port 8080, handling all requests
    log.Fatal(http.ListenAndServe(":8080", nil))
}
