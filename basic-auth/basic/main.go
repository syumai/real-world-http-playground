package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	addr        = ":8080"
	defaultUser = "user"
	defaultPass = "pass"
)

func authorized(auth string) bool {
	sl := strings.SplitAfter(auth, "Basic ")
	if len(sl) != 2 {
		return false
	}

	b := sl[1]
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	_, err := base64.StdEncoding.Decode(decoded, []byte(b))
	if err != nil {
		return false
	}

	up := strings.Split(string(decoded), ":")
	user, pass := up[0], up[1]
	return user == defaultUser && pass == defaultPass
}

func handler(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !authorized(auth) {
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
