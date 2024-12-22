package graphs

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	h := &NodeHeap{}

	nodes := []*Node{
		{Distance: 5},
		{Distance: 1},
		{Distance: 2},
		{Distance: 3},
		{Distance: 4},
	}

	for _, n := range nodes {
		heap.Push(h, n)
	}

	assert.Equal(t, 1, heap.Pop(h).(*Node).Distance)
	assert.Equal(t, 2, heap.Pop(h).(*Node).Distance)
	assert.Equal(t, 3, heap.Pop(h).(*Node).Distance)
	assert.Equal(t, 4, heap.Pop(h).(*Node).Distance)
	assert.Equal(t, 5, heap.Pop(h).(*Node).Distance)
}
