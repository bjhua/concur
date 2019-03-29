package finelockprioqueue

import (
	"fmt"
	"lib"
	"sync"
)

type status int

const (
	Empty status = iota
	Changing
	Ready
)

//////////////////
type node struct{
	*sync.Mutex

	priority int
	status status
}

func NewNode(priority int)*node{
	return &node{Mutex: &sync.Mutex{}, priority: priority, status: Changing}
}

func (n *node)String()string{
	return fmt.Sprintf("%d", n.priority)
}

type PriorityQueue struct{
	*sync.RWMutex

	arr []*node  // we use [1, ..., last], that is, root is at 1
	last int     // points to the last element, initially 0
}

func NewPriorityQueue(numElements int)*PriorityQueue{
	q := &PriorityQueue{RWMutex: &sync.RWMutex{}}
	q.arr = make([]*node, numElements+1, numElements+1)
	q.last = 0
	return q
}

func (q *PriorityQueue)Insert(priority int) {
	const rootIndex = 1
	var childIndex, parentIndex int
	backOff := lib.BackoffGen(1, 512)

	fresh := NewNode(priority)
	q.Lock()
	q.last++
	index := q.last
	q.arr[index] = fresh
	if index==rootIndex{ // this is the first element
		// no need to do any adjustment, the root is ready
		fresh.Lock()
		fresh.status = Ready
		fresh.Unlock()
		q.Unlock()
		return
	}
	q.Unlock()
	//fmt.Printf("store [%d]=%d\n", index, priority)

	q.RLock()
	defer q.RUnlock()

	childIndex = index
	for childIndex>rootIndex{
		parentIndex = childIndex/2

		//fmt.Printf("parent: %d, child: %d\n", parentIndex, localIndex)
		// local spin
		parent := q.arr[parentIndex]
		child := q.arr[childIndex]
	RETRY:
		parent.Lock()
		child.Lock()
		if parent.status != Ready{
			//fmt.Printf("spin on: %d\n", parent.priority)
			child.Unlock()
			parent.Unlock()
			backOff()
			goto RETRY
		}

		if parent.priority<= child.priority {
			child.status = Ready
			child.Unlock()
			parent.Unlock()
			return
		}
		swap(child, parent)
		childIndex = parentIndex

		child.Unlock()
		parent.Unlock()
	}
	q.arr[rootIndex].status = Ready
}

func (q *PriorityQueue)RemoveMin()(int, error){

	return 0, nil
}

func swap(n1, n2 *node){
	//
	priority := n1.priority
	n1.priority = n2.priority
	n2.priority = priority
	//
	status := n1.status
	n1.status = n2.status
	n2.status = status
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





