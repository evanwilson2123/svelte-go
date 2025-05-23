package main

import (
	"log"
	"net/http"
	"stoic/handlers"
	"stoic/middleware"
)

func main() {
	// init the server
	mux := http.NewServeMux()
	/* TODO: 
		- add handlers 
	*/
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/start-reflecting", handlers.SignInHandler)

	// wrap the server in the cors middleware
	handler := middleware.CorsMiddleware(mux)
	// set the port
	addr := ":8080"
	// start the server
	log.Printf("Server starting on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}