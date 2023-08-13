package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go RunServer(&wg, 3000, MuxServerOne())
	go RunServer(&wg, 8080, MuxServerTwo())
	wg.Wait()
}

func RunServer(wg *sync.WaitGroup, port int, mux http.Handler) {
	defer wg.Done()

	fmt.Printf("Starting server 2 at port %v\n", port)
	url := "localhost:" + strconv.Itoa(port)

	err := http.ListenAndServe(url, mux)
	if err != nil {
		log.Fatal(err)
	}
}
