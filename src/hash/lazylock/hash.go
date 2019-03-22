package lazylock

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	initCaps   = 16
	loadFactor = 0.75
)

//////////////////
// the hash
type HashSet struct {
	*sync.RWMutex

	arr        []unsafe.Pointer // *coarselock
	numItems   int
	loadFactor float32
}

func New() *HashSet {
	h := &HashSet{RWMutex: &sync.RWMutex{}, loadFactor: loadFactor}
	h.arr = make([]unsafe.Pointer, initCaps, initCaps)
	// don't initialize the array elements here, instead, initialize
	// it lazily during the first insertion
	/*
	for i:=0; i<initCaps; i++{
		h.arr = append(h.arr, NewList())
	}
	*/
	return h
}

func (h *HashSet) Insert(x int) bool {
	h.RLock()
	buk := x % len(h.arr)

	value := Singleton(x)
	inserted := false
	if atomic.CompareAndSwapPointer(&h.arr[buk], nil, (unsafe.Pointer)(value)) {
		inserted = true
	} else {
		l := (*list)(h.arr[buk])
		l.Lock() // coarselock lock
		inserted = l.insert(x)
		l.Unlock() // coarselock unlock
	}
	h.RUnlock()

	if inserted {
		h.Lock()
		h.numItems++
		if h.numItems > (int)((float32)(len(h.arr))*h.loadFactor) {
			h.resize()
		}
		h.Unlock()
		return true
	}
	return false
}

func (h *HashSet) resize() {
	newLen := 2 * len(h.arr)
	arr := make([]unsafe.Pointer, newLen, newLen)

	for _, l := range h.arr {
		lx := (*list)(l)
		if lx == nil {
			continue
		}
		cur := lx.node
		for cur != nil {
			v := cur.data
			buk := v % newLen
			ly := (*list)(arr[buk])
			if ly != nil {
				ly.insert(v)
			} else {
				arr[buk] = (unsafe.Pointer)(Singleton(v))
			}
			cur = cur.next
		}
	}
	h.arr = arr
	return
}

func (h *HashSet) Lookup(x int) bool {
	defer h.RUnlock()

	h.RLock()
	buk := x % len(h.arr)
	if (h.arr[buk]) == nil {
		return false
	}
	l := (*list)(h.arr[buk])
	l.RLock()
	found := l.lookup(x)
	l.RUnlock()
	return found
}

func (h *HashSet) Delete(x int) bool {
	h.RLock()
	buk := x % len(h.arr)
	deleted := false
	if atomic.LoadPointer(&h.arr[buk]) != nil {
		l := (*list)(h.arr[buk])
		l.Lock()
		deleted = l.delete(x)
		l.Unlock()
	}
	h.RUnlock()
	h.Lock()
	defer h.Unlock()
	if deleted {
		h.numItems--
	}
	return deleted
}

func (h *HashSet) Dump() {
	nonemptyBuckets := 0
	longestList := 0
	numSingletons := 0
	for _, l := range h.arr {
		//fmt.Printf("%d ==> ", i)
		l := (*list)(l)
		if l == nil {
			//fmt.Printf("\n")
			continue
		}
		cur := l.node
		if cur != nil {
			nonemptyBuckets++
			if cur.next == nil {
				numSingletons++
			}
		}
		curList := 0
		for cur != nil {
			curList++
			//fmt.Printf("%d, ", cur.data)
			cur = cur.next
		}
		if curList > longestList {
			longestList = curList
		}
		//fmt.Printf("\n")
	}
	fmt.Printf("buckets=%d\nnonempty=%d\nlongest=%d\nsingletons=%d\n",
		len(h.arr), nonemptyBuckets, longestList, numSingletons)
	fmt.Printf("====================\n\n")
}
