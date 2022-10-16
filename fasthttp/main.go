package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

func main() {
	ln, err := reuseport.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("error in reuseport listener: %v", err)
	}

	if err = fasthttp.Serve(ln, requestHandler); err != nil {
		log.Fatalf("error in fasthttp Server: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	resp := struct {
		Message string
	}{
		Message: "Hello World",
	}

	ctx.SetContentType("application/json")

	json.NewEncoder(ctx).Encode(resp)
}
