//非并发安全的 bitmap 计算
package bitmap

import (
	"math"
)

type BitMapStatistics struct {
	bits []byte
	max  int
}

//一个byte有8位,可以用来表示8个数字，最大数/8 的整数加 1 为即存放最大数所需的容量
func NewStatistics(max ...int) *BitMapStatistics {
	_max := math.MaxInt32
	if len(max) > 0 && max[0] > 0 {
		_max = max[0]
	}

	return &BitMapStatistics{make([]byte, _max>>3+1), _max}
}

// num >> 3 求出 num 所在的位置
// num & 7 取余，知识点：a % (2^n) 等价于 a & (2^n - 1)
func (b *BitMapStatistics) Add(num int) {
	index, pos := num>>3, num&7

	b.bits[index] |= 1 << pos

	if b.max < num {
		b.max = num
	}
}

func (b *BitMapStatistics) Exists(num int) bool {
	index, pos := num>>3, num&7
	return b.bits[index]&(1<<pos) != 0
}

func (b *BitMapStatistics) Remove(num int) {
	index, pos := num>>3, num&7
	b.bits[index] &^= 1 << pos
}
