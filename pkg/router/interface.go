package router

type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	DELETE(path string, handler HandlerFunc)
	Group(prefix string, middlewares ...Middleware) *RouteGroup
	Serve(port string) error
	Use(middleware Middleware)
}

type HandlerFunc func(c *Context)

type Middleware func(HandlerFunc) HandlerFunc
