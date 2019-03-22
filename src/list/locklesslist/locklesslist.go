package locklesslist

import (
	myatomic "concur/list/atomic"
	"fmt"
	"math"
	"unsafe"
)

type node struct {
	data int32
	*myatomic.MarkablePtr
}

type Slist struct {
	head *node
	tail *node
}

func New() *Slist {
	l := &Slist{
		head: &node{data:    math.MinInt32,
			MarkablePtr: myatomic.New(nil)},
		tail: &node{data:    math.MaxInt32,
			MarkablePtr: myatomic.New(nil)}}
	l.head.MarkablePtr = myatomic.New(unsafe.Pointer(l.tail))
	return l
}

func (l *Slist) Add(x int32) bool {
	var pred, cursor *node

	for {
		pred = l.head
		_, cursorRef := pred.Get()
		cursor = (*node)(cursorRef)
		for cursor.data < x {
			pred = cursor
			_, cursorRef = cursor.Get()
			cursor = (*node)(cursorRef)
		}

		if cursor.data == x {
			return false
		}
		fresh := &node{data: x,
			MarkablePtr: myatomic.New(unsafe.Pointer(cursor))}
		if pred.CompareAndSwap(unsafe.Pointer(cursor), unsafe.Pointer(fresh), false, false) {
			return true
		}
	}
}

func (l *Slist) Delete(x int32) bool {
	var pred, cursor *node

	for {
		pred = l.head
		_, cursorRef := pred.Get()
		cursor = (*node)(cursorRef)
		for cursor.data < x {
			pred = cursor
			_, cursorRef = cursor.Get()
			cursor = (*node)(cursorRef)
		}
		if cursor.data == x {
			_, succRef := cursor.Get()
			swappedMark := cursor.CompareAndSwap(succRef, succRef, false, true)
			if !swappedMark{
				continue
			}
			// to this point, the edge has been marked
			swapped := pred.CompareAndSwap(cursorRef, succRef, false, false)
			if !swapped{
				// unmark
				if !cursor.CompareAndSwap(succRef, succRef, true, false){
					panic("impossible")
				}
				continue
			}
			return true
		}
		return false
	}
}

func (l *Slist) Find(x int32) bool {
	var pred, cursor *node

	pred = l.head
	_, cursorRef := pred.Get()
	cursor = (*node)(cursorRef)
	for cursor.data < x {
		pred = cursor
		_, cursorRef = cursor.Get()
		cursor = (*node)(cursorRef)
	}

	if cursor.data == x{
		marked, _ := cursor.Get()
		if !marked {
			return true
		}
	}
	return false
}

func (l *Slist) Dump() {
	fmt.Printf("[")
	cursor := l.head
	for cursor != nil {
		marked, c := cursor.Get()
		fmt.Printf("(%d, %t), ", cursor.data, marked)

		cursor = (*node)(c)
	}
	fmt.Printf("]\n")
}
