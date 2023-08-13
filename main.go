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
	numberOfRunningServers := 2
	port1 := 8080
	port2 := 8081
	wg.Add(numberOfRunningServers)
	go RunServer(&wg, port1, MuxServerOne())
	go RunServer(&wg, port2, MuxServerTwo())
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
