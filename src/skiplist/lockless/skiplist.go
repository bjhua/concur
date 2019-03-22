package locklessskiplist

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	MaxLevel = 12
)

//////////////////
type SkipList struct{
	*list
	numItems int32
}

func NewSkipList()*SkipList{
	h := &SkipList{list: NewList(MaxLevel)}
	return h
}

func NewSkipListSweep()*SkipList{
	h := &SkipList{list: NewList(MaxLevel)}
	go func(){
		for {
			h.Compact()
			time.Sleep(1 * time.Millisecond)
		}
	}()
	return h
}

func (h *SkipList)Insert(x int)bool {
	numLevels := genLevel()
	inserted := h.insert(x, numLevels)
	if inserted{
		atomic.AddInt32(&h.numItems, 1)
	}
	return inserted
}


func (h *SkipList)Lookup(x int)bool{
	_, _, found := h.lookup(x)
	return found
}

func (h *SkipList)Delete(x int)bool{
	deleted := h.delete(x)
	if deleted{
		atomic.AddInt32(&h.numItems, -1)
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
		if k==MaxLevel{
			break
		}
	}
	return k
}

func (h *SkipList)Dump(){
	fmt.Printf("items = %d\n", h.numItems)
	marked := false
	var pPtr unsafe.Pointer = nil
	for i:=MaxLevel-1; i>=0; i-- {
		fmt.Printf("level[%2d] ==> ", i)
		cur := h.list
		for cur != nil {

			marked, pPtr = cur.succs[i].Get()
			if marked{
				fmt.Printf("[%3d], ", cur.data)
			}else{
				fmt.Printf("%3d, ", cur.data)
			}
			if cur.succs[i]==nil {
				break
			}
			p := (*list)(pPtr)
			//_, nextPtr := cur.succs[0].Get()
			//cur = (*coarselock)(nextPtr)
			cur = p
			/*
			for cur!=p{
				fmt.Print("   , ")
				_, nextPtr = cur.succs[0].Get()
				cur = (*coarselock)(nextPtr)
			}
			*/
		}
		fmt.Printf("\n")
	}
}

func (h *SkipList)Compact(){
	//h.Dump()
	h.list.compact()

}



