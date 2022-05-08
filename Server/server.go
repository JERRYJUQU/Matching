package main

import "net/http"

type HTTPHandler struct{}

const SERVER_ADDR = "127.0.0.1:8080"

func (handler HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")

	res.Write(data)
}

func main() {
	var handler HTTPHandler

	http.ListenAndServe(SERVER_ADDR, handler)
}
