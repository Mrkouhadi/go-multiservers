package main

import (
	"net/http"
)

// Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to Home page"))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to Contact page"))
}

func SupportHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Weclcome to support page"))
}

func ShopHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to shop page"))
}
