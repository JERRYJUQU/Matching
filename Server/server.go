package main

import (
	"log"
	"net/http"
)

const SERVER_ADDR = "127.0.0.1:8080"

func main() {
	log.Print("Starting angular-bootstrap-demo on address: ", SERVER_ADDR)
	http.Handle("/", http.FileServer(http.Dir("../TheOneProject/dist/angular-bootstrap-demo")))
	http.ListenAndServe(SERVER_ADDR, nil)
}
