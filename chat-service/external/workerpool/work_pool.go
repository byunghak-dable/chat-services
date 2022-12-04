package workerpool

import (
	"sync"
)

type WorkerPool struct {
	jobChan  chan func()
	isClosed bool
}

func New(wg *sync.WaitGroup, count int) *WorkerPool {
	jobChan := make(chan func())
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			for job := range jobChan {
				job()
			}
			wg.Done()
		}()
	}
	return &WorkerPool{
		jobChan:  jobChan,
		isClosed: false,
	}
}

func (p *WorkerPool) RegisterJob(job func()) {
	if p.isClosed {
		// TODO: need logging
		return
	}
	p.jobChan <- job
}

func (p *WorkerPool) Close() {
	close(p.jobChan)
	p.isClosed = true
	// TODO: need logging
}
