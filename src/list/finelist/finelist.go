package finelist

import (
	"fmt"
	"math"
	"sync"
)

type node struct{
	lock sync.Locker

	data int32
	next *node
}
type Slist struct{
	head *node
	tail *node
}

func New()*Slist{
	l := &Slist{
		head: &node{lock: &sync.Mutex{},data: math.MinInt32},
		tail: &node{lock: &sync.Mutex{},data: math.MaxInt32}}
	l.head.next = l.tail
	return l
}

func (l *Slist)Add(x int32)bool{
	var pred, cursor *node
	defer func() {
		//cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	pred = l.head
	pred.lock.Lock()
	cursor = pred.next
	//cursor.lock.Lock()
	for cursor.data < x{
		cursor.lock.Lock()
		pred.lock.Unlock()
		pred = cursor
		cursor = cursor.next
		//cursor.lock.Lock()
	}
	if cursor.data == x{
		return false
	}
	fresh := &node{lock: &sync.Mutex{}, data: x, next: cursor}
	pred.next = fresh
	return true
}


func (l *Slist)Delete(x int32)bool{
	var pred, cursor *node
	defer func() {
		cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	pred = l.head
	pred.lock.Lock()
	cursor = pred.next
	cursor.lock.Lock()
	for cursor.data < x{
		pred.lock.Unlock()
		pred = cursor
		cursor = cursor.next
		cursor.lock.Lock()
	}
	if cursor.data == x{
		pred.next = cursor.next
		return true
	}
	return false
}


func (l *Slist)Find(x int32)bool{
	var pred, cursor *node
	defer func() {
		cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	pred = l.head
	pred.lock.Lock()
	cursor = pred.next
	cursor.lock.Lock()
	for cursor.data < x{
		pred.lock.Unlock()
		pred = cursor
		cursor = cursor.next
		cursor.lock.Lock()
	}
	if cursor.data == x{
		return true
	}
	return false
}

func (l *Slist)Dump(){
	fmt.Printf("[")
	cursor := l.head
	for cursor != nil{
		fmt.Printf("%d", cursor.data)
		if cursor.next != nil{
			fmt.Printf(", ")
		}

		cursor = cursor.next
	}
	fmt.Printf("]\n")
}