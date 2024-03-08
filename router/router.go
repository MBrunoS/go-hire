package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	r := &Router{
		mux: http.NewServeMux(),
	}
	return r
}

func (r *Router) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("GET "+path, handler)
}

func (r *Router) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("POST "+path, handler)
}

func (r *Router) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("PUT "+path, handler)
}

func (r *Router) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.mux.HandleFunc("DELETE "+path, handler)
}

func (r *Router) Serve(port string) {
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, r.mux)
}
