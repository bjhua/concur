package slist

import (
	"fmt"
	"math"
)

type node struct{
	data int32
	next *node
}
type Slist struct{
	head *node
	tail *node
}

func New()*Slist{
	l := &Slist{head: &node{data: math.MinInt32}, tail: &node{data: math.MaxInt32}}
	l.head.next = l.tail
	return l
}

func (l *Slist)Add(x int32)bool{
	pred := l.head
	cursor := pred.next
	for cursor.data < x{
		pred = cursor
		cursor = cursor.next
	}
	if cursor.data == x{
		return false
	}
	fresh := &node{data: x, next: cursor}
	pred.next = fresh
	return true
}


func (l *Slist)Delete(x int32)bool{
	pred := l.head
	cursor := pred.next
	for cursor.data < x{
		pred = cursor
		cursor = cursor.next
	}
	if cursor.data == x{
		pred.next = cursor.next
		return true
	}
	return false
}


func (l *Slist)Find(x int32)bool{
	pred := l.head
	cursor := pred.next
	for cursor.data < x{
		pred = cursor
		cursor = cursor.next
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