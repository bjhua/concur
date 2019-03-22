package coarselock

import (
	"fmt"
	"sync"
)

const (
	initCaps = 16
	loadFactor = 0.75
)

//////////////////
// the hash
type HashSet struct{
	*sync.RWMutex

	arr []*list
	numItems int
	loadFactor float32
}

func New()*HashSet{
	h := &HashSet{RWMutex: &sync.RWMutex{}, loadFactor: loadFactor}
	h.arr = make([]*list, 0, initCaps)
	for i:=0; i<initCaps; i++{
		h.arr = append(h.arr, &list{})
	}
	return h
}

func (h *HashSet)Insert(x int)bool {
	defer h.Unlock()

	h.Lock()
	buk := x % len(h.arr)
	l := h.arr[buk]
	inserted := l.insert(x)
	if inserted{
		h.numItems++
		if h.numItems > (int)((float32)(len(h.arr))*h.loadFactor){
			h.resize()
		}
		return true
	}
	return false
}

func (h *HashSet)resize(){
	newLen := 2 *len(h.arr)
	arr := make([]*list, 0, newLen)
	for i:=0; i<newLen; i++{
		arr = append(arr, &list{})
	}
	for _, l := range h.arr{
		cur := l.node
		for cur!=nil{
			v := cur.data
			buk := v % newLen
			arr[buk].insert(v)
			cur = cur.next
		}
	}
	h.arr = arr
	return
}

func (h *HashSet)Lookup(x int)bool{
	defer h.RUnlock()

	h.RLock()
	buk := x % len(h.arr)
	l := h.arr[buk]
	return l.lookup(x)
}

func (h *HashSet)Delete(x int)bool{
	defer h.Unlock()

	h.Lock()
	buk := x % len(h.arr)
	l := h.arr[buk]
	deleted := l.delete(x)
	if deleted{
		h.numItems--
	}
	return deleted
}

func (h *HashSet)Dump(){
	for i, l := range h.arr{
		fmt.Printf("%d ==> ", i)
		cur := l.node
		for cur!=nil{
			fmt.Printf("%d, ", cur.data)
			cur = cur.next
		}
		fmt.Printf("\n")
	}
}
