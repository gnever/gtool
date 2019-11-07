//go test queue* -v
package queue_test

import (
	"fmt"
	"testing"

	"github.com/gnever/gtool/queue"
)

func TestDemo(t *testing.T) {
	n := 100
	q := queue.New()
	for i := 0; i < n; i++ {
		q.Push(i)
	}

	fmt.Println("channel长度： ", q.ChannelLen())
	fmt.Println("总长度： ", q.Len())

	for {
		if v := q.Pop(); v != nil {
			fmt.Println("for got - ", v)
		}
	}
}

func TestDemoForSelect(t *testing.T) {
	n := 100
	q := queue.New()
	for i := 0; i < n; i++ {
		q.Push(i)
	}

	fmt.Println("channel长度： ", q.ChannelLen())
	fmt.Println("总长度： ", q.Len())

	for {
		select {
		case v := <-q.C:
			fmt.Println("select got - ", v)
		}
	}
}
