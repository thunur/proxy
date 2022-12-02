package main

import (
	"fmt"
	"github.com/thunur/proxy"
	"github.com/thunur/proxy/entity"
	"net/http"
	"time"
)

type Handler struct {
	proxy.Delegate
}

func (handler *Handler) BeforeRequest(entity *entity.Entity) {
	fmt.Printf("%+v", entity.GetRequestBody())
}
func (handler *Handler) BeforeResponse(entity *entity.Entity, err error) {
	fmt.Printf("%+v", entity.GetResponseBody())
}
func (handler *Handler) ErrorLog(err error) {}

func main() {
	p := proxy.NewWithDelegate(&Handler{})
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
