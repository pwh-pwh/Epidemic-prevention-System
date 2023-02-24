package task

type Pool struct {
	work chan func()   // 任务
	sem  chan struct{} // 数量
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func Default() *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, 10),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}
