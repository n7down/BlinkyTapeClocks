package queue

import (
	"sync"
)

type Queue struct {
	lock  sync.RWMutex
	items []int
}

func NewQueue() *Queue {
	q := &Queue{
		items: make([]int, 0),
	}
	return q
}

func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q Queue) Size() int {
	return len(q.items)
}

func (q Queue) Front() *int {
	q.lock.Lock()
	defer q.lock.Unlock()

	i := q.items[0]
	return &i
}

func (q Queue) Enqueue(i int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, i)
}

func (q Queue) Dequeue() *int {
	q.lock.Lock()
	defer q.lock.Unlock()

	i := q.items[0]
	q.items = q.items[1:len(q.items)]
	return &i
}
