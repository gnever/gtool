package bitmap_test

import (

	"github.com/gnever/gtool/bitmap"
	"testing"
)

func TestStatistics(T *testing.T) {
	max := 800
	b := bitmap.NewStatistics(max)
	 for i:=0; i<max; i++ {
		b.Add(i)
	}
	if b.Exists(6) != true {
		T.Error("6 not exists")
	}
	b.Remove(6)
	if b.Exists(6) == true {
		T.Error("6 remove fail")
	}
}
