package workerPool

import (
	"fmt"
	"sync/atomic"

	"github.com/gnever/gtool/list"
)

type Pool struct {
	limit  int32 // goroutine 数量限制， 默认为 10
	count  int32 // 当前运行 goroutine 数量
	closed bool  //是否关闭

	jobs *list.List //用于保存任务
}

func New(limit ...int32) *Pool {
	p := &Pool{
		limit:  10,
		count:  0,
		closed: false}

	if len(limit) > 0 && limit[0] > 1 {
		p.limit = limit[0]
	}

	p.jobs = list.New()
	return p
}

func (p *Pool) Add(function func()) error {

	if p.closed {
		return fmt.Errorf("workPool closed")
	}

	p.jobs.PushBack(function)

	cn := atomic.LoadInt32(&p.count)
	if cn >= p.limit || p.jobs.Len() == 0 {
		return nil
	}

	if atomic.CompareAndSwapInt32(&p.count, cn, cn+1) {
		p.worker()
	}

	return nil
}

func (p *Pool) worker() {
	go func() {
		var job interface{}
		defer atomic.AddInt32(&p.count, -1)
		for !p.closed {
			if job = p.jobs.PopFront(); job != nil {
				job.(func())()
			}

		}
	}()
}

func (p *Pool) Close() {
	p.closed = true
}

func (p *Pool) Cap() int32 {
	return p.limit
}

func (p *Pool) PoolCount() int32 {
	return atomic.LoadInt32(&p.count)
}

func (p *Pool) JobSizes() int {
	return p.jobs.Len()
}
