package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// 获取Header
	http.HandleFunc("/header", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.Header)
		fmt.Fprintln(rw, r.Header["Accept-Encoding"])
		fmt.Fprintln(rw, r.Header.Get("Accept-Encoding"))
	})

	// 获取Body
	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(rw, string(body))
	})

	// 获取Query
	http.HandleFunc("/query", func(rw http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		id := query["id"]
		log.Println(id)

		name := query.Get("name")
		log.Println(name)
	})

	// 获取上传文件
	http.HandleFunc("/process", func(rw http.ResponseWriter, r *http.Request) {
		// 方式1：
		// r.ParseMultipartForm(1024)
		// fileHeader := r.MultipartForm.File["uploaded"][0]
		// file, err := fileHeader.Open()

		//方式2：
		file, _, err := r.FormFile("uploaded")
		if err == nil {
			data, err := ioutil.ReadAll(file)
			if err == nil {
				fmt.Fprintln(rw, string(data))
			}
		}

	})

	server.ListenAndServe()
}
