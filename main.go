package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe("localhost:8080", nil) // go 的web 应用，内置的库可以实现= java 的mvc
}
