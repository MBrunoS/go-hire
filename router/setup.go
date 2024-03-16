package router

import (
	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/internal/infra/database/repository"
	"github.com/mbrunos/go-hire/internal/infra/http/handler"
	"github.com/mbrunos/go-hire/pkg/router"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Setup() {
	r := router.NewDefaultRouter()
	addRoutes(r)
	r.Serve(":8080")
}

// @title Go Hire API
// @version 1.0
// @description This is a simple API to manage job offers

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @BasePath /api
func addRoutes(r router.Router) {
	db := config.GetDB()
	jobRepo := repository.NewJobRepository(db)
	jobHandler := handler.NewJobHandler(jobRepo)

	r.GET("/api/jobs", jobHandler.List)
	r.POST("/api/jobs", jobHandler.Create)
	r.GET("/api/jobs/{id}", jobHandler.Get)
	r.PUT("/api/jobs/{id}", jobHandler.Update)
	r.DELETE("/api/jobs/{id}", jobHandler.Delete)

	r.GET("/swagger/{any...}", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
