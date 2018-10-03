package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	addr = ":8080"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]
	log.Printf("file_path: %s\n", filePath)

	if filePath == "" {
		filePath = "index.html"
	}

	if exists("public/" + filePath + ".br") {
		w.Header().Set("Content-Encoding", "br")
		w.Header().Set("Content-Type", "text/plain")
		filePath = filePath + ".br"
	} else if exists("public/" + filePath + ".gz") {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/plain")
		filePath = filePath + ".gz"
	}

	f, err := os.Open(fmt.Sprintf("public/%s", filePath))
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("File not found"))
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", mainHandler)
	fmt.Printf("listening localhost%s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, "../cert/localhost.pem", "../cert/localhost-key.pem", nil))
}
