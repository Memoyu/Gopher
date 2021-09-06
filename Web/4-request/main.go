package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/header", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.Header)
		fmt.Fprintln(rw, r.Header["Accept-Encoding"])
		fmt.Fprintln(rw, r.Header.Get("Accept-Encoding"))
	})

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(rw, string(body))
	})

	server.ListenAndServe()
}
