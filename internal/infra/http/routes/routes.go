package routes

import (
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/pkg/middleware"
	"github.com/mbrunos/go-hire/pkg/router"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func AddPublicRoutes(r router.Router, userHandler *handler.UserHandler, jobHandler *handler.JobHandler) {
	g := r.Group("/api")
	g.POST("/signup", userHandler.SignUp)
	g.POST("/login", userHandler.Login)
	g.GET("/jobs", jobHandler.List)
	g.GET("/jobs/{id}", jobHandler.Get)
}

func AddPrivateRoutes(r router.Router, userHandler *handler.UserHandler, jobHandler *handler.JobHandler) {
	jwt_secret := config.JWTSecret
	g := r.Group("/api", middleware.JwtAuth(jwt_secret))
	g.PUT("/users/{id}", userHandler.Update)
	g.DELETE("/users/{id}", userHandler.Delete)
	g.POST("/jobs", jobHandler.Create)
	g.PUT("/jobs/{id}", jobHandler.Update)
	g.DELETE("/jobs/{id}", jobHandler.Delete)
}

func AddSwaggerRoutes(r router.Router) {
	r.GET("/swagger/{any...}", func(c *router.Context) {
		httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		).ServeHTTP(c.Writer, c.Request)
	})
}
