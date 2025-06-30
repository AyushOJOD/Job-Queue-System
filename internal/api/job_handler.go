package api

import (
	"net/http"
	"strconv"

	"job-queue/internal/service"
	"job-queue/internal/utils"
	"job-queue/internal/worker"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	Service service.JobService
	Worker  *worker.WorkerPool
}

func NewJobHandler(s service.JobService, w *worker.WorkerPool) *JobHandler {
	return &JobHandler{
		Service: s,
		Worker:  w,
	}
}

// POST /jobs → Submit a new job
func (h *JobHandler) SubmitJob(c *gin.Context) {
	var req struct {
		Payload string `json:"payload" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.MsgPayloadRequired})
		return
	}

	// Create job
	jobID, err := h.Service.SubmitJob(req.Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.MsgFailedToSubmitJob})
		return
	}

	// Fetch job to push to queue
	job, err := h.Service.GetJobStatus(jobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.MsgFailedToFetchJob})
		return
	}

	// Add job to worker queue
	h.Worker.AddJob(job)

	c.JSON(http.StatusOK, gin.H{
		"job_id": jobID,
		"status": job.Status,
	})
}

// GET /jobs/:id → Get job status and result
func (h *JobHandler) GetJob(c *gin.Context) {
	jobID := c.Param("id")
	job, err := h.Service.GetJobStatus(jobID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": utils.MsgJobNotFound})
		return
	}

	c.JSON(http.StatusOK, job)
}

// GET /jobs → List all jobs with pagination
func (h *JobHandler) ListJobs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	jobs, err := h.Service.ListJobs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": utils.MsgFailedToFetchJobs})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
