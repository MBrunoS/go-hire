package router

import (
	"encoding/json"
	"net/http"
)

func Setup() {
	r := NewRouter()
	addRoutes(r)
	r.Serve(":8080")
}

func addRoutes(r *Router) {
	r.GET("/api/jobs", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Get all jobs")
	})

	r.POST("/api/jobs", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Create job")
	})

	r.GET("/api/jobs/{id}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Get job by id")
	})

	r.PUT("/api/jobs/{id}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Update job by id")
	})

	r.DELETE("/api/jobs/{id}", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Delete job by id")
	})

}
