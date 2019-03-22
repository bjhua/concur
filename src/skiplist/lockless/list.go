package locklessskiplist

import (
	"concur/atomic"
	"fmt"
	"math"
	"unsafe"
)

type list struct {
	data      int
	succs     []*atomic.MarkablePtr
	numLevels int //
}

func NewNode(x int, numLevels int) *list {
	n := &list{data: x, numLevels: numLevels}
	n.succs = make([]*atomic.MarkablePtr, numLevels, numLevels)
	for i:=0; i<numLevels; i++{
		n.succs[i] = atomic.NewMarkablePtr(nil)
	}
	return n
}

func NewList(maxLevels int) *list {
	tail := NewNode(math.MaxInt32, maxLevels)
	head := NewNode(math.MinInt32, maxLevels)
	for i := 0; i < maxLevels; i++ {
		head.succs[i] = atomic.NewMarkablePtr(unsafe.Pointer(tail))
	}
	return head
}

func debug(s string){

}

// methods
// Lookup some item in the given skiplist
// return false if not found
// return true if found with its predessor and successor
// Example:
/*
For the following skip:

      1,       5,       9
      1,    4, 5,       9
      1, 2, 4, 5        9
      1, 2, 4, 5, 6, 7, 9

lookup(5) will return:
[1, 1, 2, 2], [5, 5, 5, 5], true
 */
func (l *list) dumpList() {
	n := l
	for n != nil {
		marked, p := n.succs[0].Get()
		if marked{
			fmt.Printf("<%d>[%d], ", n.data, len(n.succs))
		}else{
			fmt.Printf("%d[%d], ", n.data, len(n.succs))
		}
		n = (*list)(p)
	}
	fmt.Printf("\n")
}

func dumpSlice(s string, sl []*list) {
	fmt.Printf("%s[", s)
	for _, n := range sl {
		fmt.Printf("%d, ", n.data)
	}
	fmt.Printf("]\n")
}

func (l *list) lookup(x int) ([]*list, []*list, bool) {
	var cur *list = nil
	preds := make([]*list, l.numLevels, l.numLevels)
	succs := make([]*list, l.numLevels, l.numLevels)
	pred := l

	for level := l.numLevels - 1; level >= 0; level-- {
		_, curPtr := pred.succs[level].Get()
		cur = (*list)(curPtr)

		for cur.data < x {
			pred = cur
			_, curPtr = pred.succs[level].Get()
			cur = (*list)(curPtr)
		}
		preds[level] = pred
		succs[level] = cur
	}
	marked, _ := cur.succs[0].Get()
	return preds, succs, !marked && cur.data==x
}

func (l *list) insert(x int, numLevels int) bool {
	var preds, succs []*list
	var found bool
	var pred, succ *list

RETRY:
	preds, succs, found = l.lookup(x)
	if found {
		return false
	}

	fresh := NewNode(x, numLevels)
	for level := 0; level < numLevels; level++ {
		succ = succs[level]
		fresh.succs[level] = atomic.NewMarkablePtr(unsafe.Pointer(succ))
	}
	pred = preds[0]
	succ = succs[0]
	swapped := pred.succs[0].CompareAndSwap(unsafe.Pointer(succ), unsafe.Pointer(fresh), false, false)||
		pred.succs[0].CompareAndSwap(unsafe.Pointer(succ), unsafe.Pointer(fresh), true, true)
	if !swapped{
		goto RETRY
	}
	for level := 1; level < numLevels; level++{
		for {
			pred = preds[level]
			succ = succs[level]
			//fmt.Printf("pred: %d, succ: %d\n", pred.data, succ.data)

			swapped = pred.succs[level].CompareAndSwap(unsafe.Pointer(succ), unsafe.Pointer(fresh), false, false) ||
				pred.succs[level].CompareAndSwap(unsafe.Pointer(succ), unsafe.Pointer(fresh), true, true)
			if swapped {
				//fmt.Printf("swapped node: level[%d], %d, with pred, succ: [%d, %d]\n", level, x, pred.data, succ.data)
				break
			}
			//fmt.Printf("swapped value: %t, relook for: %d\n", swapped, x)
			//l.dumpList()
			preds, succs, _ = l.lookup(x)
			//dumpSlice("preds", preds)
			//dumpSlice("succs", succs)
		}
	}
	return true
}

func (l *list) delete(x int) bool {
	_, succs, found := l.lookup(x)
	if !found {
		return false
	}
	target := succs[0]
	for level := target.numLevels-1; level>=1; level--{
		marked, ptr := target.succs[level].Get()
		for !marked{
			target.succs[level].CompareAndSwap(ptr, ptr, false, true)
			marked, ptr = target.succs[level].Get()
		}
	}
	marked, ptr := target.succs[0].Get()
	for !marked{
		swapped := target.succs[0].CompareAndSwap(ptr, ptr, false, true)
		if swapped{
			return true
		}
		marked, ptr = target.succs[0].Get()
	}
	return false
}

func (l *list)compact(){
	var pred, cur *list = nil, nil
	fmt.Printf("compact starting\n")
	defer func(){
		fmt.Printf("compact finished\n")
	}()

	RETRY:
	for level := l.numLevels - 1; level >= 0; level-- {
		pred = l
		for {
			_, curPtr := pred.succs[level].Get()
			if curPtr==nil{
				break
			}
			cur = (*list)(curPtr)
			marked, nextPtr := cur.succs[level].Get()
			if marked{
				swapped := pred.succs[level].CompareAndSwap(curPtr, nextPtr, false, false)
				if !swapped{
					goto RETRY
				}
			}else{
				pred = cur
			}
		}
	}
}


