package lockless_queue

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

type node struct {
	data int
	next unsafe.Pointer // *node
}

type Queue struct{
	head unsafe.Pointer
	tail *node
}

func New()*Queue{
	n := &node{}
	q := &Queue{head: unsafe.Pointer(n), tail: n}
	return q
}

func (q *Queue)Enq(x int) {
	n := &node{data: x}
	for {
		last := q.tail
		swapped := atomic.CompareAndSwapPointer(&last.next, nil, unsafe.Pointer(n))
		if swapped {
			q.tail = n
			return
		}
	}
}

func (q *Queue)Deq()(int, error){
	for {
		first := (*node)(q.head)
		succ := (*node)(first.next)
		if succ == nil {
			return 0, errors.New("empty queue")
		}
		swapped := atomic.CompareAndSwapPointer(&q.head, unsafe.Pointer(first), unsafe.Pointer(succ))
		if swapped {
			return succ.data, nil
		}
	}
}
