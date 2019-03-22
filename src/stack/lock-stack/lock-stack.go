package elim_stack

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

type node struct {
	data int
	next *node
}

type Stack struct{
	top unsafe.Pointer // *node
}

func New()*Stack{
	return &Stack{}
}

func (q *Stack)Push(x int) {
	n := &node{data: x}
	for {
		first := (*node)(q.top)
		n.next = first
		swapped := atomic.CompareAndSwapPointer(&q.top, unsafe.Pointer(first), unsafe.Pointer(n))
		if swapped {
			return
		}
		q.backoff()
	}
}

func (q *Stack)Pop()(int, error){
	for {
		first := (*node)(q.top)
		if first==nil{
			return 0, errors.New("empty stack")
		}
		succ := first.next
		swapped := atomic.CompareAndSwapPointer(&q.top, unsafe.Pointer(first), unsafe.Pointer(succ))
		if swapped {
			return first.data, nil
		}
		q.backoff()
	}
}

func (q *Stack)backoff(){
	//nop
}

