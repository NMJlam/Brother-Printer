package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "os/exec"
)

const (
	PRINTER_NAME = "brother_ql.700"
)

func printHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the json is malformed
	var request Request
	if malformedJson(&request, w, r) {
		return
	}

	// Valdate the json
	if !validJson(&request, w) {
		return
	}

	if err := formatLabel(
		request.ItemId,
		request.Serial,
		request.Name,
	); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Ok:    false,
			Error: "Json is malformed",
		})
	}

	// cmd := exec.Command("lp", "-d", PRINTER_NAME, "-n", fmt.Sprintf("%d", &request.Quantity), "temp/label.png")
	// err := cmd.Run()
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(ErrorResponse{
	// 		Ok:    false,
	// 		Error: err.Error(),
	// 	})
	// 	return
	// }

	json.NewEncoder(w).Encode(SuccessResponse{
		Ok:     true,
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
