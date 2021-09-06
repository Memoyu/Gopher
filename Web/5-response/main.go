package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(rw http.ResponseWriter, r *http.Request) {
	str := `<html>  
	<head>
	<title>Go Web</title>
</head>
<body>
	Hello World
</body>
	<html>`
	rw.Write([]byte(str)) // response 响应内容
}

func writeHeaderExample(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(501) // writeHeader用于写入响应的状态码，默认相应为 http.StatusOK
	fmt.Fprintln(rw, "No Such Service")
}

func setHeaderExample(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Location", "http://google.com") // 设置Header必须在WriteHeader之前。
	rw.WriteHeader(501)                              // writeHeader之后是无法修改Header
	fmt.Fprintln(rw, "No Such Service")
}

type Post struct {
	User    string
	Threads []string
}

func jsonExample(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "",
		Threads: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	rw.Write(json)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/setheader", setHeaderExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
