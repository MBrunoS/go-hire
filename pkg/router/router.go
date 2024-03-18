package router

import (
	"fmt"
	"net/http"
)

type DefaultRouter struct {
	mux *http.ServeMux
}

func NewDefaultRouter() Router {
	r := &DefaultRouter{
		mux: http.NewServeMux(),
	}
	return r
}

func (r *DefaultRouter) GET(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc("GET "+path, handler)
}

func (r *DefaultRouter) POST(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc("POST "+path, handler)
}

func (r *DefaultRouter) PUT(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc("PUT "+path, handler)
}

func (r *DefaultRouter) DELETE(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc("DELETE "+path, handler)
}

func (r *DefaultRouter) Serve(port string) error {
	fmt.Println("Server is running on port :" + port)
	return http.ListenAndServe(":"+port, r.mux)
}

func (r *DefaultRouter) Handle(route string, handler http.HandlerFunc) {
	r.mux.HandleFunc(route, handler)
}
