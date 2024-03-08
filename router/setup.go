package router

import (
	"github.com/mbrunos/go-hire/handler"
)

func Setup() {
	r := NewRouter()
	addRoutes(r)
	r.Serve(":8080")
}

func addRoutes(r *Router) {
	r.GET("/api/jobs", handler.GetAllJobs)
	r.POST("/api/jobs", handler.CreateJob)
	r.GET("/api/jobs/{id}", handler.GetJob)
	r.PUT("/api/jobs/{id}", handler.UpdateJob)
	r.DELETE("/api/jobs/{id}", handler.DeleteJob)
}
