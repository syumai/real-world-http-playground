package http

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

var handlers = map[string]func(*ResponseWriter, *Request){}

type Request struct {
	Method, Path string
	Header       Header
	Body         []byte
}

type Header map[string]string

func (h Header) Get(key string) string {
	s, ok := h[strings.ToLower(key)]
	if !ok {
		return ""
	}
	return s
}

type ResponseWriter struct {
	net.Conn
	Status int
	Header Header
}

func NewResponseWriter(conn net.Conn) *ResponseWriter {
	return &ResponseWriter{
		Conn:   conn,
		Status: 200,
		Header: Header{
			"Content-Type": "text/html; charset=utf-8",
		},
	}
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.Status = status
}

func (w *ResponseWriter) Write(p []byte) (int, error) {
	var totalLength int

	// Write first line
	n, err := fmt.Fprintf(w.Conn, "HTTP/1.0 %d %s\n", w.Status, StatusText(w.Status))
	if err != nil {
		return 0, err
	}
	totalLength += n

	// Write headers
	if len(p) > 0 {
		w.Header["Content-Length"] = strconv.Itoa(len(p) + 1)
	}
	for k, v := range w.Header {
		line := fmt.Sprintf("%s: %s\n", k, v)
		n, err := fmt.Fprintf(w.Conn, line)
		if err != nil {
			return 0, err
		}
		totalLength += n
	}

	// Write body
	n, err = w.Conn.Write([]byte("\n"))
	if err != nil {
		return 0, err
	}
	totalLength += n
	n, err = w.Conn.Write(p)
	if err != nil {
		return 0, err
	}
	totalLength += n

	// Write return code
	n, err = w.Conn.Write([]byte("\n"))
	if err != nil {
		return 0, err
	}
	totalLength += n
	return totalLength, nil
}

func parseRequest(conn net.Conn) (*Request, error) {
	req := &Request{
		Header: Header{},
	}
	r := bufio.NewReader(conn)

	// Parse method and path
	line, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}
	l := strings.Split(string(line), " ")
	req.Method = l[0]
	req.Path = l[1]

	for {
		line, _, err := r.ReadLine()
		if len(line) == 0 {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		header := strings.Split(string(line), ": ")
		req.Header[strings.ToLower(header[0])] = header[1]
	}

	// Return if a request does not have content body
	length, ok := req.Header["content-length"]
	if !ok {
		return req, nil
	}

	// Read body
	n, err := strconv.Atoi(length)
	if err != nil {
		return nil, err
	}

	body := make([]byte, n)
	_, err = r.Read(body)
	if err != nil {
		return nil, err
	}

	req.Body = body

	return req, nil
}

func HandleFunc(path string, handler func(*ResponseWriter, *Request)) {
	handlers[path] = handler
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	req, err := parseRequest(conn)
	if err != nil {
		return err
	}
	w := NewResponseWriter(conn)
	handler, ok := handlers[req.Path]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("Resource not found"))
	}
	handler(w, req)
	return nil
}

func ListenAndServe(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("%#v", err)
		}
		err = handleConnection(conn)
		if err != nil {
			log.Printf("%#v", err)
		}
	}
}
