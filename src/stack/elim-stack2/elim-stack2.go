package elim_stack2

import (
	myatomic "concur/atomic"
	"errors"
	"math/rand"
	"sync/atomic"
	"unsafe"
)

const maxCaps = 100

type node struct {
	data int
	next *node
}

type Stack struct{
	*myatomic.ElimArray

	top unsafe.Pointer // *node
}

func New()*Stack{
	return &Stack{ElimArray: myatomic.NewElimArray(maxCaps)}
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
		_, err := q.backoff(myatomic.Role_Yang, x)
		if err==nil{
			return
		}
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
		value, err := q.backoff(myatomic.Role_Yin, 0)
		if err==nil{
			return value, nil
		}
	}
}

func (q *Stack)backoff(role myatomic.Role, value int)(int, error){
	index := rand.Intn(q.Len())
	return q.Exchange(index, role, value)
}