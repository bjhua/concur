package optlist

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

func validate(head, tail, pred, cursor *node)bool{
	var p *node
	if pred.next != cursor{
		return false
	}
	for p=head; p!=tail; p=p.next{
		if p == pred || p.data > pred.data{
			break
		}
	}
	if p != pred{
		return false
	}
	return true
}

func (l *Slist)Add(x int32)bool{
	var pred, cursor *node
	defer func() {
		cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	for {
		pred = l.head
		cursor = pred.next
		for cursor.data < x {
			pred = cursor
			cursor = cursor.next
		}
		pred.lock.Lock()
		cursor.lock.Lock()
		// validate
		if !validate(l.head, l.tail, pred, cursor) {
			cursor.lock.Unlock()
			pred.lock.Unlock()
			continue
		}

		if cursor.data == x {
			return false
		}
		fresh := &node{lock: &sync.Mutex{}, data: x, next: cursor}
		pred.next = fresh
		return true
	}
}


func (l *Slist)Delete(x int32)bool{
	var pred, cursor *node
	defer func() {
		cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	for {
		pred = l.head
		cursor = pred.next
		for cursor.data < x {
			pred = cursor
			cursor = cursor.next
		}
		pred.lock.Lock()
		cursor.lock.Lock()
		if !validate(l.head, l.tail, pred, cursor) {
			cursor.lock.Unlock()
			pred.lock.Unlock()
			continue
		}

		if cursor.data == x {
			pred.next = cursor.next
			return true
		}
		return false
	}
}


func (l *Slist)Find(x int32)bool{
	var pred, cursor *node
	defer func() {
		cursor.lock.Unlock()
		pred.lock.Unlock()
	}()

	for {
		pred = l.head
		cursor = pred.next
		for cursor.data < x {
			pred = cursor
			cursor = cursor.next
		}
		cursor.lock.Lock()
		pred.lock.Lock()
		if !validate(l.head, l.tail, pred, cursor) {
			cursor.lock.Unlock()
			pred.lock.Unlock()
			continue
		}

		if cursor.data == x {
			return true
		}
		return false
	}
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