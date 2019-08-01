package queue

import (
	"errors"
	"sync"
)

type Queue struct {
	lock  sync.Mutex
	items []interface{}
}

func NewQueue() *Queue {
	q := &Queue{
		lock: sync.Mutex{},
	}
	return q
}

func (q Queue) Len() int {
	return len(q.items)
}

func (q Queue) Peek() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	l := len(q.items)
	if l == 0 {
		return nil, errors.New("queue is empty")
	}
	i := q.items[0]
	return i, nil
}

func (q Queue) Put(i interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.items = append(q.items, i)
}

func (q Queue) Dequeue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	l := len(q.items)
	if l == 0 {
		return nil, errors.New("queue is empty")
	}
	i := q.items[0]
	q.items = q.items[:l]
	return i, nil
}
