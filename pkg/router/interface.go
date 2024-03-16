package router

import (
	"net/http"
)

type Router interface {
	GET(path string, handler func(w http.ResponseWriter, r *http.Request))
	POST(path string, handler func(w http.ResponseWriter, r *http.Request))
	PUT(path string, handler func(w http.ResponseWriter, r *http.Request))
	DELETE(path string, handler func(w http.ResponseWriter, r *http.Request))
	Serve(port string)
}
