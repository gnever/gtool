// go test -v -count=1  pool.go  workerPool_test.go

package workerPool_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gnever/gtool/workerPool"
)

func TestNew(t *testing.T) {
	num := 10000
	tf := time.NewTicker(time.Second * 1)

	w := workerPool.New(10)

	for i := 0; i < num; i++ {
		go func(n int) {
			w.Add(func() {
				time.Sleep(time.Second * 1)
			})
		}(i)
	}

	for {
		select {
		case <-tf.C:
			fmt.Printf("cap %d\n", w.Cap())
			fmt.Printf("PoolCount %d\n", w.PoolCount())
			fmt.Printf("JobSizes %d\n", w.JobSizes())
			fmt.Println("========================")
		}
	}

}
