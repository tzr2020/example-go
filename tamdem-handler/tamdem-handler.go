package main

import (
	"log"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func logger(h http.Handler) http.Handler {
	// 将处理器函数转化为处理器后返回
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("处理器被调用: %T\n", h)
		// todo...
		h.ServeHTTP(w, r)
	})
}

// 串联处理器演示
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := helloHandler{}
	http.Handle("/hello", logger(hello)) // 串联处理器
	server.ListenAndServe()
}
