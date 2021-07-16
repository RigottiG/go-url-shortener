package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	port    int
	baseURL string
)

type Headers map[string]string

func responseWith(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
}

func extractUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseWith(w, http.StatusMethodNotAllowed, Headers{"Allow": "POST"})
		return
	}
}

func init() {
	port = 8888
	baseURL = fmt.Sprintf("http://localhost:%d", port)
}

func main() {
	http.HandleFunc("/api/shortener", Shortener)
	// http.HandleFunc("/r/", Redirect)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port), nil,
	))
}
