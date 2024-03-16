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

func (r *DefaultRouter) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("GET "+path, handler)
}

func (r *DefaultRouter) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("POST "+path, handler)
}

func (r *DefaultRouter) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("PUT "+path, handler)
}

func (r *DefaultRouter) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("DELETE "+path, handler)
}

func (r *DefaultRouter) Serve(port string) {
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, r.mux)
}
