package main

import (
	"log"
	"net/http"
	"os"
)

// Types
type APIServer struct {
	address string
	cert    string
	key     string
}

func NewAPIServer(hostAndPort string, cert string, key string) *APIServer {
	return &APIServer{
		address: hostAndPort,
		cert:    cert,
		key:     key,
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

	host := httpServer.address
	noTLS := os.Getenv("NO_TLS") == "true"

	if noTLS {
		host = ":8080"
	}

	server := http.Server{
		Addr:    host,
		Handler: midlewareChained(proxyRouter),
	}

	if noTLS {
		log.Printf("*** Serve starting without TLS ***")
		log.Printf("Sever has stated %s", host)
		return server.ListenAndServe()
	}

	log.Printf("Sever has stated %s", host)

	// Init server,create some keys to use TLS
	return server.ListenAndServeTLS(httpServer.cert, httpServer.key)
}
