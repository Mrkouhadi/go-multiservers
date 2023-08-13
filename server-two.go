package main

import "net/http"

func MuxServerTwo() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/shop", ShopHandler)
	mux.HandleFunc("/support", SupportHandler)

	return mux
}
