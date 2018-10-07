package main

import (
	"net/http"
)

const addr = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Cookie"]; !ok {
		w.Header().Add("Set-Cookie", "visit=true")
		w.Write([]byte("This is your first visit to this page."))
		return
	}
	w.Write([]byte("This is not your first visit to this page."))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
