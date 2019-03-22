package lazylist

import (
	"fmt"
	"math"
	"sync"
)

type node struct {
	lock    sync.Locker
	deleted bool

	data int32
	next *node
}

type Slist struct {
	head *node
	tail *node
}

func New() *Slist {
	l := &Slist{
		head: &node{lock: &sync.Mutex{},
			deleted: false,
			data:    math.MinInt32},
		tail: &node{lock: &sync.Mutex{},
			deleted: false,
			data:    math.MaxInt32}}
	l.head.next = l.tail
	return l
}

func validate(pred, cursor *node) bool {
	if !pred.deleted && !cursor.deleted && pred.next == cursor {
		return true
	}
	return false
}

func (l *Slist) Add(x int32) bool {
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
		if !validate(pred, cursor) {
			cursor.lock.Unlock()
			pred.lock.Unlock()
			continue
		}

		if cursor.data == x {
			return false
		}
		fresh := &node{lock: &sync.Mutex{},
			data: x,
			next: cursor}
		pred.next = fresh
		return true
	}
}

func (l *Slist) Delete(x int32) bool {
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
		if !validate(pred, cursor) {
			cursor.lock.Unlock()
			pred.lock.Unlock()
			continue
		}

		if cursor.data == x {
			cursor.deleted = true
			pred.next = cursor.next
			return true
		}
		return false
	}
}

func (l *Slist) Find(x int32) bool {
	var pred, cursor *node

	pred = l.head
	cursor = pred.next
	for cursor.data < x {
		pred = cursor
		cursor = cursor.next
	}

	if !cursor.deleted && cursor.data == x {
		return true
	}
	return false
}

func (l *Slist) Dump() {
	fmt.Printf("[")
	cursor := l.head
	for cursor != nil {
		fmt.Printf("%d", cursor.data)
		if cursor.next != nil {
			fmt.Printf(", ")
		}

		cursor = cursor.next
	}
	fmt.Printf("]\n")
}
