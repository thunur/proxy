package main

import (
	"github.com/thunur/proxy"
	"net/http"
	"time"
)

func main() {
	p := proxy.New()
	server := &http.Server{
		Addr: ":8889",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			p.ServerHandler(rw, req)
		}),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
