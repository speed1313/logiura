package main

import (
	"log"

	"github.com/speed1313/logiura/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
