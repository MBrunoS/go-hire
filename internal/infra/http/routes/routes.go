package routes

import (
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/pkg/middleware"
	"github.com/mbrunos/go-hire/pkg/router"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func addPublicRoutes(r router.Router, userHandler *handler.UserHandler, jobHandler *handler.JobHandler) {
	g := r.Group("/api")
	g.POST("/signup", userHandler.SignUp)
	g.POST("/login", userHandler.Login)
	g.GET("/jobs", jobHandler.List)
	g.GET("/jobs/{id}", jobHandler.Get)
}

func addPrivateRoutes(r router.Router, userHandler *handler.UserHandler, jobHandler *handler.JobHandler) {
	jwt_secret := config.JWTSecret
	g := r.Group("/api", middleware.JwtAuth(jwt_secret))
	g.PUT("/users/{id}", userHandler.Update)
	g.DELETE("/users/{id}", userHandler.Delete)
	g.POST("/jobs", jobHandler.Create)
	g.PUT("/jobs/{id}", jobHandler.Update)
	g.DELETE("/jobs/{id}", jobHandler.Delete)
}

func addSwaggerRoutes(r router.Router) {
	r.GET("/swagger/{any...}", func(c *router.Context) {
		httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		).ServeHTTP(c.Writer, c.Request)
	})
}

func Setup(r router.Router, userHandler *handler.UserHandler, jobHandler *handler.JobHandler) {
	addPublicRoutes(r, userHandler, jobHandler)
	addPrivateRoutes(r, userHandler, jobHandler)
	addSwaggerRoutes(r)
}
