package main
// func printHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
//
// 	// Check if the json is malformed
// 	var request Request
// 	if malformedJson(&request, w, r) {
// 		return
// 	}
//
// 	// Valdate the json
// 	if !validJson(&request, w) {
// 		return
// 	}
//
// 	json.NewEncoder(w).Encode(SuccessResponse{
// 		Ok:     true,
// 		ItemId: request.ItemId,
// 	})
// }
//
// http.HandleFunc("/printer", printHandler)
//
// fmt.Println("Server starting on http://localhost:6767")
//
// if err := http.ListenAndServe(":6768", nil); err != nil {
// 	fmt.Printf("Error starting server: %s\n", err)
// }

func main() {
	// Create blank label with specified dimensions
	formatLabel()
}
