// go test workerPool_bench_test.go pool.go  -bench="Ben.*" -benchmem -cpuprofile profile_cpu.out
// go tool pprof -pdf profile_cpu.out > profile_cpu.pdf
package workerPoolB

import (
	"testing"
	"time"

	"github.com/gnever/gtool/workerPool"
)

func BenchmarkWorkerPool(b *testing.B) {
	b.ResetTimer()

	w := workerPool.New(1000)
	for i := 0; i < b.N; i++ {
		w.Add(func() {
			time.Sleep(time.Millisecond * 1)
		})
	}
}

func BenchmarkGoroutine(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		go func() {
			time.Sleep(time.Millisecond * 1)
		}()
	}
}

func TB() {
	for i := 0; i <= 1000; i++ {
	}
}
