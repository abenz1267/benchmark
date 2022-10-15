package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Message string
		}{
			Message: "Hello World",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8080", nil)
}
