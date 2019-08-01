package queue

import (
	"errors"
)

// FIXME: add sync logic
type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	q := &Queue{}
	return q
}

func (q Queue) Peek() (interface{}, error) {
	l := len(q.items)
	if l == 0 {
		return nil, errors.New("queue is empty")
	}
	i := q.items[0]
	return i, nil
}

func (q Queue) Put(i interface{}) {
	q.items = append(q.items, i)
}

func (q Queue) Dequeue() (interface{}, error) {
	l := len(q.items)
	if l == 0 {
		return nil, errors.New("queue is empty")
	}
	i := q.items[0]
	q.items = q.items[:l-1]
	return i, nil
}
