package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danielmerchant/info344-in-class/middleware/handlers"
	"github.com/danielmerchant/info344-in-class/middleware/middleware"
)

const defaultAddr = ":80"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/time", handlers.TimeHandler)

	//TODO: wrap the mux with the Logger middleware
	loggedMux := middleware.NewLogger(mux)

	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, loggedMux))
}
