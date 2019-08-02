package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueShouldEnqueueItemsWhenCallingEnqueue(t *testing.T) {
	assert := assert.New(t)
	q := NewQueue()

	q.Enqueue(1)
	assert.Equal(1, q.Size(), "Len should be 1")

	q.Enqueue(2)
	assert.Equal(2, q.Size(), "Len should be 2")

	//assert.Fail("Not implemented")
}

func TestQueueShouldRemoveItemsWhenCallingDequeue(t *testing.T) {
	assert := assert.New(t)
	assert.Fail("Not implemented")
}

func TestQueueShouldContainItemsWhenCallingFront(t *testing.T) {
	assert := assert.New(t)
	assert.Fail("Not implemented")
}
