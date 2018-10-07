package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const (
	addr           = ":8080"
	accessCountKey = "access_count"
)

const htmlTpl = `<!doctype html>
<html>
	<body>
		<div>You've been to this page for {{.Count}} times.</div>
		<div>
			<button onclick="window.location='/';">
				Reload
			</button>
			<form method="POST" action="/reset">
				<button>Reset</button>
			</form>
		</div>
	</body>
</html>`

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.New("html").Parse(htmlTpl)
	if err != nil {
		log.Fatalf("%#v", err)
	}
}

func setCountCookie(w http.ResponseWriter, cnt int) {
	c := strconv.Itoa(cnt)
	http.SetCookie(w, &http.Cookie{
		Name:     accessCountKey,
		Value:    c,
		HttpOnly: true,
	})
}

func render(w http.ResponseWriter, cnt int) error {
	return tmpl.Execute(w, struct{ Count int }{Count: cnt})
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("resource not found"))
	}

	var cnt int

	if ck, err := r.Cookie(accessCountKey); err == nil {
		cnt, _ = strconv.Atoi(ck.Value)
	}

	cnt++
	setCountCookie(w, cnt)

	err := render(w, cnt)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
	}
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
	}

	setCountCookie(w, 0)

	err := render(w, 0)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
	}
}

func main() {
	http.HandleFunc("/", countHandler)
	http.HandleFunc("/reset", resetHandler)
	fmt.Printf("listening localhost%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
