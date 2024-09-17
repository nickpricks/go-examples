package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Message string
	Data    interface{}
	Status  int16
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welocome to homepage")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	response := APIResponse{Message: "Welcome to the homepage", Data: nil, Status: http.StatusOK}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	const port = ":9000"
	http.HandleFunc("/", home)
	fmt.Printf("Server started at port http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
