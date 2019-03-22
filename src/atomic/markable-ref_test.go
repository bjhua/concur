package atomic

import (
	"fmt"
	"testing"
	"unsafe"
)

type node struct {
	data int
	*MarkablePtr
}

func newNode(data int, next *node)*node{
	return &node{data: data, MarkablePtr: New(unsafe.Pointer(next))}
}

func (this *node)dump(){
	l := this
	fmt.Printf("[")
	for l!=nil{
		mark, ref := l.Get()
		fmt.Printf("(%d, %t), ", l.data, mark)
		l = (*node)(ref)
	}
	fmt.Printf("]\n")
}

func TestMarkableRef(t *testing.T) {
	l := newNode(1, newNode(2, newNode(3, nil)))
	l.dump()
	_, oldRef := l.Get()
	//
	newRef := newNode(4, (*node)(oldRef))
	l.CompareAndSwap(unsafe.Pointer(oldRef), unsafe.Pointer(newRef), false, false)
	l.dump()
}


