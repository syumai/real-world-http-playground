package main

import (
	"net/http"
)

const addr = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
