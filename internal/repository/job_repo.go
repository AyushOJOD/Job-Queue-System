package repository

import (
	"job-queue/internal/models"

	"gorm.io/gorm"
)

type JobRepository interface {
	CreateJob(job *models.Job) error
	GetJobByID(id string) (*models.Job, error)
	UpdateJob(job *models.Job) error
	ListJobs(limit, offset int) ([]*models.Job, error)
}

type jobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

func (r *jobRepository) CreateJob(job *models.Job) error {
	return r.db.Create(job).Error
}

func (r *jobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	if err := r.db.First(&job, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *jobRepository) UpdateJob(job *models.Job) error {
	return r.db.Save(job).Error
}

func (r *jobRepository) ListJobs(limit, offset int) ([]*models.Job, error) {
	var jobs []*models.Job
	if err := r.db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}
