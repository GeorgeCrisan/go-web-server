package main

import (
	"log"
	"net/http"
)

// Types
type APIServer struct {
	address string
}

func NewAPIServer(hostAndPort string) *APIServer {
	return &APIServer{
		address: hostAndPort,
	}
}

// Run is the method to initiate the server
func (httpServer *APIServer) Run() error {
	router := http.NewServeMux()

	// Define the middleware we want to use for this router
	midlewareChained := ChainMiddleware(
		RequestLogger,
		RequestValidator,
	)

	// Set the route functionality
	router.HandleFunc("GET /user/{userId}/token/{token}", userHandler)

	// Allows to point certain routes to other routers
	proxyRouter := http.NewServeMux()
	proxyRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    httpServer.address,
		Handler: midlewareChained(proxyRouter),
	}

	log.Printf("Sever has stated %s", httpServer.address)

	// Init server,create some keys to use TLS
	return server.ListenAndServe()
}
