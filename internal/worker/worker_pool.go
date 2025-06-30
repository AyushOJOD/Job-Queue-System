package worker

import (
	"fmt"
	"job-queue/internal/models"
	"job-queue/internal/service"
	"job-queue/internal/utils"
	"log"
)

type WorkerPool struct {
	JobChan   chan *models.Job
	NumWorker int
	Service   service.JobService
}

func NewWorkerPool(num int, service service.JobService) *WorkerPool {
	return &WorkerPool{
		JobChan:   make(chan *models.Job, utils.DefaultQueueSize),
		NumWorker: num,
		Service:   service,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.NumWorker; i++ {
		go wp.worker(i)
	}
	log.Printf(utils.MsgWorkersStarted, wp.NumWorker)
}

func (wp *WorkerPool) worker(id int) {
	for job := range wp.JobChan {
		log.Printf(utils.MsgWorkerProcessing, id, job.ID)

		err := wp.Service.ProcessJob(job)
		if err != nil {
			log.Printf(utils.MsgWorkerFailed, id, job.ID, err)
			continue
		}

		log.Printf(utils.MsgWorkerCompleted, id, job.ID)
	}
}

func (wp *WorkerPool) AddJob(job *models.Job) {
	fmt.Printf(utils.MsgJobAddedToQueue, job.ID)
	wp.JobChan <- job
}
