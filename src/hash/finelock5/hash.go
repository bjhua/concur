package finelock_hash3

import (
	"fmt"
	"sync"
)

const (
	initCaps = 2048
	lockCaps = 8
	loadFactor = 0.75
)

//////////////////
// the hash
type HashSet struct{
	lockArr *[]*sync.RWMutex
	arr []*list
	numItems int
	loadFactor float32
}

func New()*HashSet{
	h := &HashSet{loadFactor: loadFactor}
	lockArr := new([]*sync.RWMutex)
	*lockArr = make([]*sync.RWMutex, lockCaps, lockCaps)
	h.lockArr = lockArr
	for i:=0; i<lockCaps; i++{
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
func (h *HashSet)lockAll(lockArr *[]*sync.RWMutex){
	for {
		lockArr := h.lockArr

		for i:=0; i<len(*lockArr); i++{
			(*lockArr)[i].Lock()
		}
		if h.lockArr==lockArr {
			return
		}
		//panic("")

			for i:=0; i<len(*lockArr); i++{
				(*lockArr)[i].Unlock()
			}
			}
}

func (h *HashSet)unlockAll(lockArr *[]*sync.RWMutex){
	for i:=0; i<len(*lockArr); i++{
		(*lockArr)[i].Unlock()
	}
}

func (h *HashSet)rlockList(buk int){
	buk = buk%lockCaps
	for{
		lockArr := h.lockArr
		(*lockArr)[buk].RLock()
		if h.lockArr==lockArr{
			return
		}
		(*lockArr)[buk].RUnlock()
		lockArr = h.lockArr
	}
}

func (h *HashSet)runlockList(buk int){
	buk = buk%lockCaps
	lockArr := h.lockArr
	(*lockArr)[buk].RUnlock()
}

func (h *HashSet)lockList(buk int){
	buk = buk%lockCaps
	for{
		lockArr := h.lockArr
		(*lockArr)[buk].Lock()
		if h.lockArr==lockArr{
			return
		}
		(*lockArr)[buk].Unlock()
		lockArr = h.lockArr
	}
}

func (h *HashSet)unlockList(buk int){
	buk = buk%lockCaps
	lockArr := h.lockArr
	(*lockArr)[buk].Unlock()
}

func (h *HashSet)Insert(x int)bool {
	//h.RLock()
	buk := x % len(h.arr)
	h.lockList(buk)
	l := h.arr[buk]
	//l.Lock()   // coarselock lock
	inserted := l.insert(x)
	h.unlockList(buk) // coarselock unlock
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
	/*
	lockArr := new([]*sync.RWMutex)
	*lockArr = make([]*sync.RWMutex, newLen, newLen)
	for i:=0; i<newLen; i++{
		(*lockArr)[i] = &sync.RWMutex{}
	}
	*/
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
	//h.lockArr = lockArr
	h.arr = arr
	return
}

func (h *HashSet)Lookup(x int)bool{
	buk := x % len(h.arr)
	l := h.arr[buk]
	//lockArr := h.lockArr
	h.rlockList(buk)
	found := l.lookup(x)
	h.runlockList(buk)
	return found
}

func (h *HashSet)Delete(x int)bool{
	buk := x % len(h.arr)
	l := h.arr[buk]
	h.lockList(buk)
	deleted := l.delete(x)
	h.unlockList(buk)

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
	fmt.Printf("buckets=%d\nnonempty=%d\nlongest=%d\nloadFactor=%f\nsingletons=%d\nitems=%d\n",
		len(h.arr), nonemptyBuckets, longestList, (float32)(h.numItems)/(float32)(len(h.arr)), numSingletons, h.numItems)
	fmt.Printf("====================\n\n")
}
