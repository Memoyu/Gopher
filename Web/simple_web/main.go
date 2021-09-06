package main

import "net/http"

type myHandler struct{}

// 自定义handler,用于处理http请求
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	mh := myHandler{}
	server := http.Server{
		Addr:    "localhost:8080", // 监听地址
		Handler: &mh,              // Handler
	}
	server.ListenAndServe() // 开始监听网络请求
	// http.ListenAndServe("localhost:8080", nil) // 等价于上方的写法，但是没有上面的写法灵活
}
