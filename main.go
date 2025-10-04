package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/skip2/go-qrcode"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the json is malformed
	var request Request
	if malformedJson(&request, w, r) {
		return
	}
	
	// Valdate the json
	if validJson(&request, w) {
		return
	}

		

	json.NewEncoder(w).Encode(SuccessResponse{
		Ok:    true,
		ItemId: request.ItemId,
	})
}

func main() {
	http.HandleFunc("/printer", printHandler)

	fmt.Println("Server starting on http://localhost:6767")

	if err := http.ListenAndServe(":6767", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
