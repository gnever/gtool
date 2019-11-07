//go test list* -v
package list_test

import (
	"gtool/list"
	"testing"
)

func TestDemo(t *testing.T) {

	n := 10

	l := list.New()
	for i := 0; i < n; i++ {
		l.PushFront(i)
	}

	if l.Len() != n {
		t.Error("长度不一致")
	}

	l.Clear()
	if l.Len() != 0 {
		t.Error("清除失败")
	}

	for i := 0; i < n; i++ {
		l.PushFront(i)
	}

	l.PopBack()
	l.PopBack()
	l.PopBack()

	if v := l.PopBack(); v != 3 {
		t.Errorf("获取值与预期不一致, expect 3, got %d", v)
	}

}
