package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type XRequest struct {
	Content string `json:"Content"`
}

type XResponse struct {
	Data struct {
		ID string `json:"id"`
	}
	Message string `json:"message"`
}

func enableCors(w http.ResponseWriter, r *http.Request) {
	// Fix the typos: "controll" -> "Control"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func sendJSONResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := XResponse{
		Message: message,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func PostRequest(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers first, before any response
	enableCors(w, r)

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set response content type
	w.Header().Set("Content-Type", "application/json")

	// Read and parse request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var requestData XRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Msg := requestData.Content

	fmt.Printf("Received content: %s\n", requestData.Content)

	err = MakePost(requestData.Content)
	if err != nil {
		fmt.Println("Failed to post twt", err)
		sendJSONResponse(w, "Failed to post")
		return
	}

	// Send JSON response back to frontend
	sendJSONResponse(w, "Post received successfully")
}
