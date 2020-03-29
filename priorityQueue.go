/*
    a priority queue implementation, to learn from the heap impl in golang
*/
package priorityQueue

import(
"container/heap"
"fmt"
)

type Item struct{
value interface{} // an arbitrary value to track
priority int // the priority
index int // the index of the item in the heap
}

// priority queue impl heap methods and holds slice of item pointers
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i,j int){
pq[i], pq[j] = pq[j], pq[i]
pq[i].index = i
pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
n:=len(*pq)
item:=x.(*Item)
item.index = n
*pq = append(*pq, item)
}
func(pq *PriorityQueue) Pop() interface{}{
old:=*pq
n:=len(old)
item:=old[n-1]
old[n-1]=nil // zero indexed
item.index=-1 // invalidating structure
*pq = old[0:n-1] // shortening the slice
return item
}
