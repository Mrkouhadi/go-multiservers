package main

import "net/http"

func MuxServerOne() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", HomeHandler)
	mux.HandleFunc("/contact", ContactHandler)

	return mux
}
