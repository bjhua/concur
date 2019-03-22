package coarseskiplist

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	initCaps = 12
)

//////////////////
type SkipList struct{
	*sync.RWMutex

	arr []*list
	numItems int
}

func NewSkipList()*SkipList{
	h := &SkipList{RWMutex: &sync.RWMutex{}}
	h.arr = make([]*list, initCaps, initCaps)
	for i:=0; i<initCaps; i++{
		h.arr[i] = NewList()
		if i!=0{
			h.arr[i].node.down = h.arr[i-1].node
			h.arr[i].node.next.down = h.arr[i-1].node.next
		}
	}
	return h
}

func (h *SkipList)Insert(x int)bool {
	defer h.Unlock()

	h.Lock()
	maxLevel := genLevel()
	top := len(h.arr)-1
	l := h.arr[top]
	inserted := l.insert(x, maxLevel)
	if inserted{
		h.numItems++
	}
	return inserted
}

func (h *SkipList)Lookup(x int)bool{
	defer h.RUnlock()

	h.RLock()
	level := len(h.arr)-1
	found := h.arr[level].lookupRaw(x)
	return found
}

func (h *SkipList)Delete(x int)bool{
	defer h.Unlock()

	h.Lock()
	level := len(h.arr)-1
	deleted := h.arr[level].delete(x)
	if deleted{
		h.numItems--
		return true
	}
	return false
}

func genLevel()int{
	k := 1
	for{
		if rand.Float64()>=0.5{
			break
		}
		k++
		if k==initCaps{
			break
		}
	}
	return k
}

func (h *SkipList)Dump(){
	for i:= len(h.arr)-1; i>=0; i--{
		fmt.Printf("level[%d] ==> ", i)
		cur := h.arr[i].node
		for cur!=nil{
			fmt.Printf("%d, ", cur.data.data)
			cur = cur.next
		}
		fmt.Printf("\n")
	}
}
