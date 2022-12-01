package pool

import "sync"

type WorkerPool struct {
	jobChan chan func()
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
		jobChan: jobChan,
	}
}

func (p *WorkerPool) RegisterJob(job func()) {
	p.jobChan <- job
}

func (p *WorkerPool) Close() {
	close(p.jobChan)
}
