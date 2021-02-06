package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Please visit http://127.0.0.1:8080/ or http://127.0.0.1:8080/index or http://127.0.0.1:8080/about")

	httpServeByStruct()
	// httpServeByStruct2()
	// httpServeByStruct3()
	// httpServeByStruct4()
	// httpServeByStruct5()
	// httpServeByFunc()
	// httpServeByFunc2()
}

type MyHandler struct{}
type MyHandler2 struct{}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Index"))
	if err != nil {
		log.Println(err)
	}
}

func (mh2 *MyHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("About"))
	if err != nil {
		log.Println(err)
	}
}

func httpServeByStruct() {
	mh := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &mh,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func httpServeByStruct2() {
	mh := MyHandler{}
	mh2 := MyHandler2{}

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/index", &mh)
	mux.Handle("/about", &mh2)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func httpServeByStruct3() {
	mh := MyHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.DefaultServeMux.Handle("/index", &mh)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func httpServeByStruct4() {
	mh := MyHandler{}

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.Handle("/index", &mh)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func httpServeByStruct5() {
	mh := MyHandler{}

	http.Handle("/index", &mh)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Index"))
	if err != nil {
		log.Println(err)
	}
}

func httpServeByFunc() {
	http.Handle("/index", http.HandlerFunc(index))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

func httpServeByFunc2() {
	http.HandleFunc("/index", index)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
