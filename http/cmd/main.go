package main

import (
	"github.com/syumai/real-world-http-playground/http"
)

func handler(w *http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080")
}
