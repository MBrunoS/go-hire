package router

import (
	"github.com/mbrunos/go-hire/handler"

	_ "github.com/mbrunos/go-hire/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Setup() {
	r := NewRouter()
	handler.Init()
	addRoutes(r)
	r.Serve(":8080")
}

// @title Go Hire API
// @version 1.0
// @description This is a simple API to manage job offers

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @BasePath /api
func addRoutes(r *Router) {
	r.GET("/api/jobs", handler.GetAllJobs)
	r.POST("/api/jobs", handler.CreateJob)
	r.GET("/api/jobs/{id}", handler.GetJob)
	r.PUT("/api/jobs/{id}", handler.UpdateJob)
	r.DELETE("/api/jobs/{id}", handler.DeleteJob)

	r.GET("/swagger/{any...}", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
