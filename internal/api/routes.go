package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, h *JobHandler) {
	api := r.Group("/api")
	{
		api.POST("/jobs", h.SubmitJob)
		api.GET("/jobs/:id", h.GetJob)
		api.GET("/jobs", h.ListJobs)
	}
}
