package queue

import (
	"math"

	"github.com/gnever/gtool/list"
)

type Queue struct {
	list          *list.List
	C             chan interface{} //队列 channel
	closed        bool             // 是否已经关闭
	prefetchCount int              // 每次从 list 中预读多少数据到 channel 中
	trigger       chan int8        // 触发从 list 中读数据到 channel
}

const defaultQueueSize = 1000

func New(prefetchCount ...int) *Queue {
	pc := defaultQueueSize
	if len(prefetchCount) > 0 && prefetchCount[0] > 1 {
		pc = prefetchCount[0]
	}

	q := &Queue{
		list:          list.New(),
		C:             make(chan interface{}, defaultQueueSize),
		prefetchCount: pc,
		trigger:       make(chan int8, math.MaxInt32),
		closed:        false}
	go q.async()
	return q
}

func (q *Queue) async() {
	defer func() {
		q.Close()
	}()

	var num int
	var listLen int

	for !q.closed {
		//主要避免 list 为空时不停的获取 list 信息
		<-q.trigger

		for !q.closed {
			listLen = q.list.Len()
			if listLen <= 0 {
				break
			}

			if q.prefetchCount > listLen {
				num = listLen
			} else {
				num = q.prefetchCount
			}

			for _, v := range q.list.PopFronts(num) {
				q.C <- v
			}
		}

		for i := 0; i < len(q.trigger)-1; i++ {
			<-q.trigger
		}

	}

}

func (q *Queue) Pop() (v interface{}) {
	return <-q.C
}

func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
	if len(q.trigger) < defaultQueueSize {
		q.trigger <- 1
	}
}

func (q *Queue) Close() {
	q.closed = true
	q.list.Clear()
}

func (q *Queue) Len() int {
	return q.list.Len() + len(q.C)
}

func (q *Queue) ChannelLen() int {
	return len(q.C)
}
