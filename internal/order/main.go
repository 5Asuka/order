package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Starting server... Listening on :8082")

	//创建相应页面
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v\n", r.RequestURI)
		_, _ = w.Write([]byte("pong"))
	})

	//启动页面 错误处理
	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal(err)
	}
}
