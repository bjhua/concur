package finelock_hash2

import "sync"

type node struct {
	data int
	next *node
}

type list struct{
	*sync.RWMutex

	node *node
}

func NewList()*list{
	return &list{RWMutex: &sync.RWMutex{}}
}

// methods
func (l *list)lookup(x int)bool{
	cur := l.node
	for cur!=nil{
		if cur.data==x{
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *list)insert(x int)bool{
	if l.lookup(x){
		return false
	}
	n := &node{data: x, next: l.node}
	l.node = n
	return true
}

func (l *list)delete(x int)bool{
	prev := &(l.node)
	cur := l.node
	for cur!=nil && cur.data!=x{
		prev = &cur.next
		cur = cur.next
	}
	if cur!=nil{
		*prev = cur.next
		return true
	}
	return false
}

