package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		resp := struct {
			Message string
		}{
			Message: "Hello World",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe("localhost:8080", nil)
}
