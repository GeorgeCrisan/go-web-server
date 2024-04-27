package main

import (
	"log"
	"net/http"
	"strings"
)

// Types
// type Middleware func(next http.Handler) http.HandlerFunc
type Middleware func(http.Handler) http.HandlerFunc

// Chain middleware util function

func ChainMiddleware(middlwareList ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		// iterate from all middleware, invoke each middleware and pass object to next middleware
		for i := len(middlwareList) - 1; i >= 0; i-- {
			next = middlwareList[i](next)
		}

		return next.ServeHTTP
	}
}

/* Midleware definition */

// Logger middleware
func RequestLogger(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Log info method: %s, path : %s, info: %s ", req.Method, req.URL.Path, req.UserAgent())

		next.ServeHTTP(res, req)
	}
}

// Mock dumb validator as middleware proof of concept
func RequestValidator(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		token := "1234"

		if !strings.Contains(req.URL.Path, token) {
			log.Printf("Token is required !!!")
			http.Error(res, "Unauthorised so fuck off", http.StatusUnauthorized)
			// Early return
			return
		}

		log.Printf("Security Token found, move on")

		next.ServeHTTP(res, req)
	}
}
