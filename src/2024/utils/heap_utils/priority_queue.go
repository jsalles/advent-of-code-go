package pqueue

import "container/heap"

// Item holds the value and its priority
type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

type priorityQueue[T any] []*Item[T]

// Required heap interface methods
func (pq priorityQueue[T]) Len() int           { return len(pq) }
func (pq priorityQueue[T]) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq priorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// PriorityQueue is the public interface
type PriorityQueue[T any] struct {
	items *priorityQueue[T]
}

// New creates a new priority queue
func New[T any]() *PriorityQueue[T] {
	items := make(priorityQueue[T], 0)
	return &PriorityQueue[T]{
		items: &items,
	}
}

// Push adds an item to the queue
func (pq *PriorityQueue[T]) Push(value T, priority int) {
	item := &Item[T]{
		Value:    value,
		Priority: priority,
	}
	heap.Push(pq.items, item)
}

// Pop removes and returns the highest priority item
func (pq *PriorityQueue[T]) Pop() *Item[T] {
	if pq.items.Len() == 0 {
		return nil
	}
	return heap.Pop(pq.items).(*Item[T])
}

// Peek returns the highest priority item without removing it
func (pq *PriorityQueue[T]) Peek() *Item[T] {
	if pq.items.Len() == 0 {
		return nil
	}
	return (*pq.items)[0]
}

// Len returns the number of items in the queue
func (pq *PriorityQueue[T]) Len() int {
	return pq.items.Len()
}

// IsEmpty returns true if the queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.items.Len() == 0
}

// Clear removes all items from the queue
func (pq *PriorityQueue[T]) Clear() {
	*pq.items = make(priorityQueue[T], 0)
}
