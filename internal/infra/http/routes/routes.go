package routes

import (
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/pkg/router"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func AddUserRoutes(r router.Router, userHandler *handler.UserHandler) {
	g := r.Group("/api/users")
	g.POST("/", userHandler.Create)
	g.PUT("/{id}", userHandler.Update)
	g.DELETE("/{id}", userHandler.Delete)
}

func AddJobRoutes(r router.Router, jobHandler *handler.JobHandler) {
	g := r.Group("/api/jobs")
	g.GET("/", jobHandler.List)
	g.POST("/", jobHandler.Create)
	g.GET("/{id}", jobHandler.Get)
	g.PUT("/{id}", jobHandler.Update)
	g.DELETE("/{id}", jobHandler.Delete)
}

func AddSwaggerRoutes(r router.Router) {
	r.GET("/swagger/{any...}", func(c *router.Context) {
		httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		).ServeHTTP(c.Writer, c.Request)
	})
}
