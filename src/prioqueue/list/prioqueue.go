package listqueue

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

//////////////////
type node struct{
	*sync.Mutex

	priority int
	deleted bool
	next *node
}

func NewNode(priority int)*node{
	return &node{Mutex: &sync.Mutex{}, priority: priority}
}

type PriorityQueue struct{
	head *node
	tail *node
}

func NewPriorityQueue()*PriorityQueue{
	q := &PriorityQueue{head: NewNode(math.MinInt32), tail:NewNode(math.MaxInt32)}
	q.head.next = q.tail
	return q
}

func validate(pred, cur *node)bool{
	return !pred.deleted && pred.next == cur
}

func (q *PriorityQueue)Insert(priority int) {
	var pred, cur *node
	defer func() {
		pred.Unlock()
		cur.Unlock()
	}()

	RETRY:
	pred = q.head
	cur = pred.next
	for cur.priority < priority{
		pred = cur
		cur = cur.next
	}
	pred.Lock()
	cur.Lock()
	if !validate(pred, cur){
		pred.Unlock()
		cur.Unlock()
		goto RETRY
	}
	fresh := NewNode(priority)
	fresh.next = cur
	pred.next = fresh
	return
}

func (q *PriorityQueue)RemoveMin()(int, error){
	var pred, cur *node
	defer func(){
		pred.Unlock()
		cur.Unlock()
	}()

	RETRY:
	pred = q.head
	cur = pred.next
	pred.Lock()
	cur.Lock()
	if !validate(pred, cur){
		pred.Unlock()
		cur.Unlock()
		goto RETRY
	}

	if cur==q.tail{
		return 0, errors.New("empty priority queue")
	}
	pred.next = cur.next
	return cur.priority, nil
}

func (q *PriorityQueue)Dump(name string){
	cur := q.head
	for cur !=nil{
		fmt.Printf("%d, ", cur.priority)
		cur = cur.next
	}
	fmt.Printf("\n")
}
