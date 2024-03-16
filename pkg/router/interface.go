package router

import (
	"net/http"
)

type Router interface {
	GET(path string, handler http.HandlerFunc)
	POST(path string, handler http.HandlerFunc)
	PUT(path string, handler http.HandlerFunc)
	DELETE(path string, handler http.HandlerFunc)
	Serve(port string) error
	GetHandler() http.Handler
	Handle(route string, handler http.HandlerFunc)
}
