package requestlogger

import (
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/pkg/router"
)

func Logger(next router.HandlerFunc) router.HandlerFunc {
	logger := config.GetLogger()

	return func(c *router.Context) {
		logger.Info(c.Request.Method, c.Request.URL.Path)
		next(c)
	}
}
