package lazyskiplist

import (
	"fmt"
	"math/rand"
	"sync/atomic"
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

func (h *SkipList)Insert(x int)bool {
	numLevels := genLevel()
	inserted := h.insert(x, numLevels)
	if inserted{
		atomic.AddInt32(&h.numItems, 1)
	}
	return inserted
}


func (h *SkipList)Lookup(x int)bool{
	_, _, topLevel := h.lookup(x, false)
	if topLevel==-1{
		return false
	}
	return true
}

func (h *SkipList)Delete(tid int, x int)bool{
	deleted := h.delete(tid, x)
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
	for i:=MaxLevel-1; i>=0; i-- {
		fmt.Printf("level[%d] ==> ", i)
		cur := h.list
		for cur != nil {
			if len(cur.succs)>i{
				fmt.Printf("%d, ", cur.data)
			}else{
				fmt.Printf("  , ")
			}
			cur = cur.succs[i]
		}
		fmt.Printf("\n")
	}
}
