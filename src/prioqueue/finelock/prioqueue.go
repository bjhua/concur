package pathlockprioqueue

import (
	"concur/lib"
	"fmt"
	"sync"
)

//////////////////
type node struct{
	*sync.Mutex

	priority int
}

func NewNode(priority int)*node{
	return &node{Mutex: &sync.Mutex{}, priority: priority}
}

func (n *node)String()string{
	return fmt.Sprintf("%d", n.priority)
}

type PriorityQueue struct{
	*sync.Mutex

	arr []*node  // we use [1, ..., last], that is, root is at 1
	last int     // points to the last element, initially 0
}

func NewPrioQueue(numElements int)*PriorityQueue{
	q := &PriorityQueue{Mutex: &sync.Mutex{}}
	q.arr = make([]*node, numElements+1, numElements+1)
	return q
}

func (q *PriorityQueue)Insert(priority int) {
	const rootIndex = 1
	fresh := NewNode(priority)
	q.Lock()
	q.last++
	index := q.last
	q.arr[index] = fresh
	if index==rootIndex{
		// no need to do any adjustment
		q.Unlock()
		return
	}
	q.Unlock()
	//fmt.Printf("store [%d]=%d\n", index, priority)

	childIndex := index
	parentIndex := childIndex/2
	for parentIndex>0{
		//fmt.Printf("parent: %d, child: %d\n", parentIndex, localIndex)
		parent := q.arr[parentIndex]
		child := q.arr[childIndex]
		parent.Lock()
		child.Lock()

		if parent.priority<= child.priority {
			child.Unlock()
			parent.Unlock()
			return
		}
		tmp := parent.priority
		parent.priority = child.priority
		child.priority = tmp
		childIndex = parentIndex
		parentIndex = childIndex/2

		child.Unlock()
		parent.Unlock()
	}
}

func (q *PriorityQueue)Remove()int{
	return 0
}

func (q *PriorityQueue)Dump(name string){
	d := lib.NewDot(name)
	for i:=1; 2*i<=q.last; i++{
		//fmt.Printf("dumping: %d\n", i)
		parent := q.arr[i]
		leftChild := q.arr[2*i]
		d.AppendEdge(parent, leftChild)
		if 2*i+1<=q.last{
			rightChild := q.arr[2*i+1]
			d.AppendEdge(parent, rightChild)
		}
	}
	d.Layout()
}
