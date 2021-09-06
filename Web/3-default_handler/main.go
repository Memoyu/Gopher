package main

import "net/http"

// Go默认Handler
// 1、NotFoundHandler (func NotFoundHandler() Handler)
// 返回一个Handler,它给每个请求的相应都是“404 page not found”

// 2、RedirectHandler (func RedirectHandler(url string, code int) Handler)
// 返回一个handler,它把每个请求使用给定的状态码跳转到指定的url

// 3、StripPrefix (func StripPrefix(prefix string, h handler) Handler)
// 返回一个handler,它从请求Url中去掉指定的前缀，然后调用另外一个Handler.
//   -如果请求的Url与提供的Url前缀不符，则但会404
//   -有点类似于中间件，修饰了另外一个Handler

// 4、TimeoutHandler (func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler)
// 返回一个handler,用于指定在时间内运行func传入的handler
//   -相当于一个修饰器，h就是指定时间内运行的Handler（被修饰的Handler）,dt 允许h运行的时间, msg 如果超时，则返回消息

// 5、FileServer (func FileServer(root FileSystem) Handler)
// 返回一个handler, 使用基于root的文件系统来响应请求

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(rw, r, "wwwroot"+r.URL.Path)
	// })
	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot"))) // 同上，简化了代码
}
