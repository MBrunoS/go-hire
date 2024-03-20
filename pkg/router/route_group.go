package router

type RouteGroup struct {
	middlewares []Middleware
	router      *DefaultRouter
	prefix      string
}

func (rg *RouteGroup) GET(path string, handler HandlerFunc) {
	rg.router.GET(rg.prefix+path, rg.applyMiddlewares(handler))
}

func (rg *RouteGroup) POST(path string, handler HandlerFunc) {
	rg.router.POST(rg.prefix+path, rg.applyMiddlewares(handler))
}

func (rg *RouteGroup) PUT(path string, handler HandlerFunc) {
	rg.router.PUT(rg.prefix+path, rg.applyMiddlewares(handler))
}

func (rg *RouteGroup) DELETE(path string, handler HandlerFunc) {
	rg.router.DELETE(rg.prefix+path, rg.applyMiddlewares(handler))
}

func (rg *RouteGroup) applyMiddlewares(handler HandlerFunc) HandlerFunc {
	for i := len(rg.middlewares) - 1; i >= 0; i-- {
		handler = rg.middlewares[i](handler)
	}
	return handler
}
