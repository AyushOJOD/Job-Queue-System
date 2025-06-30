package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusProcessing JobStatus = "processing"
	StatusCompleted  JobStatus = "completed"
	StatusFailed     JobStatus = "failed"
)

type Job struct {
	ID        string     `json:"id" gorm:"primaryKey"`
	Payload   string     `json:"payload"`
	Status    JobStatus  `json:"status"`
	Result    string     `json:"result"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Hook to generate UUID before creating
func (job *Job) BeforeCreate(tx *gorm.DB) (err error) {
	job.ID = uuid.New().String()
	return
}
