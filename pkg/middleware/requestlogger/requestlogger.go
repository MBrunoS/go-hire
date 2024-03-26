package requestlogger

import (
	"github.com/mbrunos/go-hire/pkg/logger"
	"github.com/mbrunos/go-hire/pkg/router"
)

func RequestLogger(logger logger.Logger) router.Middleware {
	return func(next router.HandlerFunc) router.HandlerFunc {

		return func(c *router.Context) {
			logger.Info(c.Request.Method, c.Request.URL.Path)
			next(c)
		}
	}
}
