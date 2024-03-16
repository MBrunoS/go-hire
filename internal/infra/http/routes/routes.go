package routes

import (
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/pkg/router"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func AddUserRoutes(r router.Router, userHandler *handler.UserHandler) {
	r.POST("/api/users", userHandler.Create)
	r.GET("/api/users/{email}", userHandler.Get)
	r.PUT("/api/users/{id}", userHandler.Update)
	r.DELETE("/api/users/{email}", userHandler.Delete)
}

func AddJobRoutes(r router.Router, jobHandler *handler.JobHandler) {
	r.GET("/api/jobs", jobHandler.List)
	r.POST("/api/jobs", jobHandler.Create)
	r.GET("/api/jobs/{id}", jobHandler.Get)
	r.PUT("/api/jobs/{id}", jobHandler.Update)
	r.DELETE("/api/jobs/{id}", jobHandler.Delete)
}

func AddSwaggerRoutes(r router.Router) {
	r.GET("/swagger/{any...}", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
