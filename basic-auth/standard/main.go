package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	addr        = ":8080"
	defaultUser = "user"
	defaultPass = "pass"
)

func handler(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != defaultUser || pass != defaultPass {
		w.Header().Add("WWW-Authenticate", "Basic")
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}
	w.Write([]byte("Authorized"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("listening localhost%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
