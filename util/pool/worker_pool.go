package pool

import "runtime"

type Pool struct {
	workers []worker
	size    int
}

type worker struct {
	channel chan interface{}
}

func NewPool(size int, f func(interface{})) (p Pool) {
	if f == nil {
		return
	}
	if size <= 0 {
		size = runtime.GOMAXPROCS(0)
	}
	p.size = size
	p.workers = make([]worker, p.size, p.size)
	for i := 0; i < size; i++ {
		p.workers[i] = initWorker(f)
	}
	return
}

func (p *Pool) Put(f interface{}) {
	i := 0
	for {
		select {
		case p.workers[i].channel <- f:
			return
		default:
			i++
			if i == p.size {
				i = 0
			}
		}
	}
}

func (p *Pool) IsRunning(f interface{}) bool {
	i := 0
	for i < p.size {
		select {
		case p.workers[i].channel <- f:
			return true
		default:
			i++
		}
	}
	return false
}

func (p *Pool) ClosePool() {
	for i := 0; i < p.size; i++ {
		p.workers[i].closeWorker()
	}
}

func initWorker(f func(interface{})) (w worker) {
	w.channel = make(chan interface{}, 1)
	go func() {
		var job interface{}
		for job = range w.channel {
			f(job)
		}
	}()
	return w
}

func (w *worker) closeWorker() {
	close(w.channel)
}
