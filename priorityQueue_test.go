package priorityQueue

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(test *testing.T) {
	items := map[int]int{1: 2, 2: 3, 3: 4}
	pq := make(PriorityQueue, len(items))
	i := 0
	for v, p := range items {
		pq[i] = &Item{value: v, priority: p, index: i}
		i++
	}
	h := &pq
	heap.Init(h)
	p := heap.Pop(h).(*Item)
	var c *Item
	for h.Len() > 0 {
		c = heap.Pop(h).(*Item)
		if p.priority > c.priority {
			test.Error("mismatched priorities")
		} else if p.value.(int) > c.value.(int) { // need type assertion as made the value a pseudo generic
			test.Error("mismatched values")
		}
		p = c
	}
}
