package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from version xx"))
	return
}

func main() {
	http.HandleFunc("/", hello)
	port := getPort()
	host := getHost()
	log.Println(fmt.Sprintf("listening on port %s:%s", host, port))
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}
	return port
}

func getHost() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	return host
}
