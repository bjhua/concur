package finelock_hash3

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
	lockArr *[]*sync.RWMutex // one lock per bucket
	arr []*list
	numItems int
	loadFactor float32
}

func New()*HashSet{
	h := &HashSet{loadFactor: loadFactor}
	t := make([]*sync.RWMutex, initCaps, initCaps)
	h.lockArr = &t
	for i:=0; i<initCaps; i++{
		(*h.lockArr)[i] = &sync.RWMutex{}
	}
	h.arr = make([]*list, initCaps, initCaps)
	for i:=0; i<initCaps; i++{
		h.arr[i] = NewList()
	}
	return h
}

// this strategy may be too slow, for we should
// have to grab too many locks one time
func (h *HashSet)lockAll(arr *[]*sync.RWMutex){
	for {
		for i:=0; i<len(*arr); i++{
			(*arr)[i].Lock()
		}
		if h.lockArr==arr {
			return
		}
		//panic("")

			for i:=0; i<len(*arr); i++{
				(*arr)[i].Unlock()
			}

			arr = h.lockArr
	}
}

func (h *HashSet)unlockAll(arr *[]*sync.RWMutex){
	for i:=0; i<len(*arr); i++{
		(*arr)[i].Unlock()
	}
}

func (h *HashSet)lockList(lockArr *[]*sync.RWMutex, buk int){
	for{
		(*lockArr)[buk].Lock()
		if h.lockArr==lockArr{
			return
		}
		(*lockArr)[buk].Unlock()
		lockArr = h.lockArr
	}
}

func (h *HashSet)unlockList(lockArr *[]*sync.RWMutex, buk int){
	(*lockArr)[buk].Unlock()
}

func (h *HashSet)Insert(x int)bool {
	//h.RLock()
	buk := x % len(h.arr)
	h.lockList(h.lockArr, buk)
	l := h.arr[buk]
	//l.Lock()   // coarselock lock
	inserted := l.insert(x)
	h.unlockList(h.lockArr, buk) // coarselock unlock
	//h.RUnlock()

	if inserted {
		h.lockAll(h.lockArr)
		lockArr := h.lockArr

		h.numItems++
		if h.numItems > (int)((float32)(len(h.arr))*h.loadFactor){
			h.resize()
		}
		h.unlockAll(lockArr)
		return true
	}
	return false
}

func (h *HashSet)resize(){
	newLen := 2 *len(h.arr)
	lockArr := new([]*sync.RWMutex)
	*lockArr = make([]*sync.RWMutex, newLen, newLen)
	for i:=0; i<newLen; i++{
		(*lockArr)[i] = &sync.RWMutex{}
	}
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
	h.lockArr = lockArr
	h.arr = arr
	return
}

func (h *HashSet)Lookup(x int)bool{
	buk := x % len(h.arr)
	l := h.arr[buk]
	lockArr := h.lockArr
	h.lockList(lockArr, buk)
	found := l.lookup(x)
	h.unlockList(h.lockArr, buk)
	return found
}

func (h *HashSet)Delete(x int)bool{
	buk := x % len(h.arr)
	l := h.arr[buk]
	h.lockList(h.lockArr, buk)
	deleted := l.delete(x)
	h.unlockList(h.lockArr, buk)

	arr := h.lockArr
	h.lockAll(arr)
	defer func(){
		h.unlockAll(h.lockArr)
	}()

	if deleted{
		fmt.Printf("--\n")
		h.numItems--
	}
	return deleted
}


func (h *HashSet)Dump(){
	nonemptyBuckets := 0
	longestList := 0
	numSingletons := 0
	for i, l := range h.arr{
		fmt.Printf("%d ==> ", i)
		l := (*list)(l)
		if l==nil{
			fmt.Printf("\n")
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
			fmt.Printf("%d, ", cur.data)
			cur = cur.next
		}
		if curList > longestList{
			longestList = curList
		}
		fmt.Printf("\n")
	}
	fmt.Printf("buckets=%d\nnonempty=%d\nlongest=%d\nloadFactor=%f\nsingletons=%d\nitems=%d\n",
		len(h.arr), nonemptyBuckets, longestList, (float32)(h.numItems)/(float32)(len(h.arr)), numSingletons, h.numItems)
	fmt.Printf("====================\n\n")
}
