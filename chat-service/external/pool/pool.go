package pool

import (
	"sync"
)

type Pool struct {
	wg    *sync.WaitGroup
	jobCh chan func()
}

func New(wg *sync.WaitGroup, buffersize int) *Pool {
	return &Pool{
		wg:    wg,
		jobCh: make(chan func(), buffersize),
	}
}

func (p *Pool) Generate(count int) {
	p.wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer p.wg.Done()
			for job := range p.jobCh {
				job()
			}
		}()
	}
}

func (p *Pool) RegisterJob(job func()) {
	p.jobCh <- job
}

func (p *Pool) Close() {
	close(p.jobCh)
}
