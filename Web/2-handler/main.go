package main

import "net/http"

// Handler 的使用

type helloHandler struct{}

// 自定义handler,用于处理http请求
func (m *helloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello world"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("About"))
}

func welcome(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Welcom"))
}

func main() {
	hh := helloHandler{}
	ah := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080", // 监听地址
		Handler: nil,              // Handler
	}
	// 采用Handle注册自定义handler
	http.Handle("/hello", &hh)
	http.Handle("/about", &ah)

	// 采用HandleFunc形式进行handler注册。（底层仍然是调用handle()）
	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Home!"))
	})
	http.HandleFunc("/welcome", welcome)
	http.Handle("/hi", http.HandlerFunc(welcome)) // 与上相同，使用http.HandlerFunc(welcome)构造ServeHTTP

	server.ListenAndServe() // 开始监听网络请求
	// http.ListenAndServe("localhost:8080", nil) // 等价于上方的写法，但是没有上面的写法灵活
}
