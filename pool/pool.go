package pool

import (
	"fmt"
)

type Pool interface {
	Schedule(task func()) error
}

type pool struct {
	work chan func()
	sem  chan struct{}
}

func New(size int) Pool {
	return &pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}
func (p *pool) Schedule(task func()) error {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		fmt.Println(len(p.sem))
		go p.worker(task)
	}
	return nil
}
func (p *pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}
