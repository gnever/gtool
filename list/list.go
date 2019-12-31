package list

import (
	"container/list"
	"sync"
)

type List struct {
	mu   *sync.RWMutex
	list *list.List
}

func New() *List {
	return &List{
		mu:   new(sync.RWMutex),
		list: list.New()}
}

func (l *List) PushFront(v interface{}) (e *list.Element) {
	l.mu.Lock()
	e = l.list.PushFront(v)
	l.mu.Unlock()
	return
}

func (l *List) PushBack(v interface{}) (e *list.Element) {
	l.mu.Lock()
	e = l.list.PushBack(v)
	l.mu.Unlock()
	return
}

func (l *List) PopFront() (v interface{}) {
	l.mu.Lock()
	if e := l.list.Front(); e != nil {
		v = l.list.Remove(e)
	}
	l.mu.Unlock()
	return
}

func (l *List) PopBack() (v interface{}) {
	l.mu.Lock()
	if e := l.list.Back(); e != nil {
		v = l.list.Remove(e)
	}
	l.mu.Unlock()
	return
}

func (l *List) PopFronts(num int) (v []interface{}) {
	l.mu.Lock()

	len := l.list.Len()
	if len == 0 {
		l.mu.Unlock()
		return
	}

	if num > len {
		num = len
	}

	v = make([]interface{}, num)
	for i := 0; i < num; i++ {
		v[i] = l.list.Remove(l.list.Front())
	}

	l.mu.Unlock()
	return
}

func (l *List) PopBacks(num int) (v []interface{}) {
	l.mu.Lock()

	len := l.list.Len()
	if len == 0 {
		l.mu.Unlock()
		return
	}

	if num > len {
		num = len
	}

	v = make([]interface{}, num)
	for i := 0; i < num; i++ {
		v[i] = l.list.Remove(l.list.Back())
	}

	l.mu.Unlock()
	return
}

func (l *List) Len() (len int) {
	l.mu.RLock()
	len = l.list.Len()
	l.mu.RUnlock()
	return
}

func (l *List) Clear() {
	l.mu.Lock()
	l.list = list.New()
	l.mu.Unlock()
}
