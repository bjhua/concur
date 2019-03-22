package finelock_hash3

import (
	"fmt"
)

const (
	initCaps = 16
	loadFactor = 0.75
)

//////////////////
// the hash
type HashSet struct{
	//*sync.RWMutex

	arr []*list
	numItems int
	loadFactor float32
}

func New()*HashSet{
	h := &HashSet{loadFactor: loadFactor}
	h.arr = make([]*list, initCaps, initCaps)
	for i:=0; i<initCaps; i++{
		h.arr[i] = NewList()
	}
	return h
}

// this strategy may be too slow, for we should
// have to grab too many locks one time
func (h *HashSet)lockAll(arr []*list){
	for _, l := range arr{
		l.Lock()
	}
}

func (h *HashSet)unlockAll(arr []*list){
	for _, l := range arr{
		l.Unlock()
	}
}

func (h *HashSet)Insert(x int)bool {
	//h.RLock()
	buk := x % len(h.arr)
	l := h.arr[buk]
	l.Lock()   // coarselock lock
	inserted := l.insert(x)
	l.Unlock() // coarselock unlock
	//h.RUnlock()

	if inserted {
		arr := h.arr
		h.lockAll(arr)
		h.numItems++
		if h.numItems > (int)((float32)(len(h.arr))*h.loadFactor){
			h.resize()
		}
		h.unlockAll(arr)
		return true
	}
	return false
}

func (h *HashSet)resize(){
	newLen := 2 *len(h.arr)
	arr := make([]*list, newLen, newLen)
	for i:=0; i<newLen; i++{
		arr[i] = NewList()
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
	buk := x % len(h.arr)
	l := h.arr[buk]
	l.RLock()
	found := l.lookup(x)
	l.RUnlock()
	return found
}

func (h *HashSet)Delete(x int)bool{
	buk := x % len(h.arr)
	l := h.arr[buk]
	l.Lock()
	deleted := l.delete(x)
	l.Unlock()

	arr := h.arr
	h.lockAll(arr)
	defer h.unlockAll(arr)
	if deleted{
		h.numItems--
	}
	return deleted
}


func (h *HashSet)Dump(){
	nonemptyBuckets := 0
	longestList := 0
	numSingletons := 0
	for _, l := range h.arr{
		//fmt.Printf("%d ==> ", i)
		l := (*list)(l)
		if l==nil{
			//fmt.Printf("\n")
			continue
		}
		cur := l.node
		if cur!=nil{
			nonemptyBuckets++
			if cur.next==nil{
				numSingletons++
			}
		}
		curList := 0
		for cur!=nil{
			curList++
			//fmt.Printf("%d, ", cur.data)
			cur = cur.next
		}
		if curList > longestList{
			longestList = curList
		}
		//fmt.Printf("\n")
	}
	fmt.Printf("buckets=%d\nnonempty=%d\nlongest=%d\nloadFactor=%f\nsingletons=%d\n",
		len(h.arr), nonemptyBuckets, longestList, (float32)(h.numItems)/(float32)(len(h.arr)), numSingletons)
	fmt.Printf("====================\n\n")
}
