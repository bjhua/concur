package coarselockprioqueue

import (
	"concur/lib"
	"fmt"
	"sync"
)

//////////////////
type node struct{
	//*sync.Mutex
	priority int
}

func NewNode(priority int)*node{
	return &node{priority: priority}
}

func (n *node)String()string{
	return fmt.Sprintf("%d", n.priority)
}

type PriorityQueue struct{
	*sync.Mutex

	arr []*node // int -> *node
	last int
}

func NewPrioQueue(maxElems int)*PriorityQueue{
	q := &PriorityQueue{Mutex: &sync.Mutex{}}
	q.arr = make([]*node, maxElems+1, maxElems+1)
	return q
}

func (q *PriorityQueue)Insert(priority int) {
	q.Lock()
	defer q.Unlock()

	q.last++
	last := q.last
	n := NewNode(priority)
	q.arr[last] = n
	//fmt.Printf("store [%d]=%d\n", index, priority)

	// adjust all parent
	localIndex := last
	parentIndex := localIndex/2
	for parentIndex>0 && parentIndex!=localIndex{
		//fmt.Printf("parent: %d, child: %d\n", parentIndex, localIndex)
		current := q.arr[localIndex]
		parent := q.arr[parentIndex]

		if parent.priority <= current.priority {
			return
		}
		tmp := parent.priority
		parent.priority = current.priority
		current.priority = tmp

		localIndex = parentIndex
		parentIndex = localIndex/2
	}
}

func (q *PriorityQueue)RemoveMin()int{
	return 0
}

func (q *PriorityQueue)Dump(name string){
	var left, right *node

	d := lib.NewDot(name)
	for i:=1; 2*i<=q.last; i++{
		//fmt.Printf("dumping: %d\n", i)
		parent := q.arr[i]
		left = q.arr[2*i]
		d.AppendEdge(parent, left)
		if 2*i+1<=q.last{
			right = q.arr[2*i+1]
			d.AppendEdge(parent, right)
		}
	}
	d.Layout()
}
