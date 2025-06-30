package service

import (
	"fmt"
	"job-queue/internal/models"
	"job-queue/internal/repository"
	"job-queue/internal/utils"
	"time"
)

type JobService interface {
	SubmitJob(payload string) (string, error)
	GetJobStatus(id string) (*models.Job, error)
	ListJobs(limit, offset int) ([]*models.Job, error)
	ProcessJob(job *models.Job) error
}

type jobService struct {
	repo repository.JobRepository
}

func NewJobService(repo repository.JobRepository) JobService {
	return &jobService{repo: repo}
}

// SubmitJob creates a new job with "pending" status
func (s *jobService) SubmitJob(payload string) (string, error) {
	job := &models.Job{
		Payload: payload,
		Status:  models.StatusPending,
	}

	if err := s.repo.CreateJob(job); err != nil {
		return "", err
	}
	return job.ID, nil
}

// GetJobStatus returns job details by ID
func (s *jobService) GetJobStatus(id string) (*models.Job, error) {
	return s.repo.GetJobByID(id)
}

// ListJobs with pagination
func (s *jobService) ListJobs(limit, offset int) ([]*models.Job, error) {
	return s.repo.ListJobs(limit, offset)
}

// ProcessJob simulates job execution
func (s *jobService) ProcessJob(job *models.Job) error {
	// Update status to processing
	job.Status = models.StatusProcessing
	job.UpdatedAt = time.Now()
	if err := s.repo.UpdateJob(job); err != nil {
		return err
	}

	// Simulate work (replace with real logic if needed)
	result := fmt.Sprintf(utils.MsgJobProcessedResult, job.Payload)

	// Update job to completed
	job.Status = models.StatusCompleted
	job.Result = result
	job.UpdatedAt = time.Now()

	return s.repo.UpdateJob(job)
}
