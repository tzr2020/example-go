package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// 解析模板源
		// t, _ := template.ParseFiles("tmpl.html")

		// 相当于创建模板，然后解析文本文件内容到模板
		// t := template.New("tmpl.html")
		// t, _ = t.ParseFiles("tmpl.html")

		// 解析模板源，增加解析模板源时的错误处理
		t := template.Must(template.ParseFiles("tmpl.html"))

		// 执行模板
		t.Execute(rw, "Hello, world!")
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
