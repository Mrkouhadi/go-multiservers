package main

import (
	"net/http"
)

// Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to Home page on port:8080"))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to Contact page on port:3000"))
}

func SupportHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to support page on port:5000"))
}

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to shop page on port:5500"))
}
