package router

import (
	"fmt"
	"net/http"
)

type DefaultRouter struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func NewDefaultRouter() Router {
	return &DefaultRouter{
		mux: http.NewServeMux(),
	}
}

func (r *DefaultRouter) GET(path string, handler HandlerFunc) {
	handlerWithMiddlewares := r.applyMiddlewares(handler)

	httpHandler := func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		handlerWithMiddlewares(ctx)
	}

	r.mux.HandleFunc("GET "+path, httpHandler)
}

func (r *DefaultRouter) POST(path string, handler HandlerFunc) {
	handlerWithMiddlewares := r.applyMiddlewares(handler)

	httpHandler := func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		handlerWithMiddlewares(ctx)
	}

	r.mux.HandleFunc("POST "+path, httpHandler)
}

func (r *DefaultRouter) PUT(path string, handler HandlerFunc) {
	handlerWithMiddlewares := r.applyMiddlewares(handler)

	httpHandler := func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		handlerWithMiddlewares(ctx)
	}

	r.mux.HandleFunc("PUT "+path, httpHandler)
}

func (r *DefaultRouter) DELETE(path string, handler HandlerFunc) {
	handlerWithMiddlewares := r.applyMiddlewares(handler)

	httpHandler := func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		handlerWithMiddlewares(ctx)
	}

	r.mux.HandleFunc("DELETE "+path, httpHandler)
}

func (r *DefaultRouter) Group(prefix string, middlewares ...Middleware) *RouteGroup {
	return NewRouteGroup(r, prefix, middlewares...)
}

func (r *DefaultRouter) Serve(port string) error {
	fmt.Println("Server is running on port :" + port)
	return http.ListenAndServe(":"+port, r.mux)
}

func (r *DefaultRouter) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *DefaultRouter) applyMiddlewares(handler HandlerFunc) HandlerFunc {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	return handler
}
