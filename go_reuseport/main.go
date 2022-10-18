package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	reuseport "github.com/kavu/go_reuseport"
)

func main() {
	runtime.GOMAXPROCS(1)

	listener, err := reuseport.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	server := &http.Server{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Message string
		}{
			Message: "Hello World",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(resp)
	})

	panic(server.Serve(listener))
}
