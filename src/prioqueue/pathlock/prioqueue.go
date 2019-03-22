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
	arr []*node
	last int

	*sync.Mutex
}

func NewPrioQueue(numItems int)*PriorityQueue{
	q := &PriorityQueue{Mutex: &sync.Mutex{}}
	q.arr = make([]*node, numItems+1, numItems+1)
	return q
}

func (q *PriorityQueue)lockAll(n int){
	var cur *node = nil
	//path := make([]*node, 0, 1)

	for n!=0{
		//fmt.Printf("looping: %d\n", n)
		cur = q.arr[n]
		for cur==nil{
			//fmt.Printf("nesting looping: %d\n", n)
			cur = q.arr[n]
		}
		//fmt.Printf("lock: %d\n", n)
		cur.Lock()
		//path = append(path, cur)

		n = n/2
	}
	//return path
}

func (q *PriorityQueue)unlockAll(path []*node){
	for i:=len(path)-1; i>=0; i-- {
		//fmt.Printf("unlock: %d\n", n)
		cur := path[i]
		cur.Unlock()
	}
}

func (q *PriorityQueue)unlockAll2(n int){
	var cur *node = nil
	//path := make([]*node, 0, 1)

	for n!=0{
		//fmt.Printf("looping: %d\n", n)
		cur = q.arr[n]
		//fmt.Printf("lock: %d\n", n)
		cur.Unlock()
		//path = append(path, cur)

		n = n/2
	}
}


func (q *PriorityQueue)Insert(priority int) {
	q.Lock()
	q.last++
	index := q.last
	//fmt.Printf("insert starting at pos[%d]: %d\n", index, priority)

	cur := NewNode(priority)
	q.arr[index] = cur
	q.Unlock()

	q.lockAll(index)
	// adjust all parent
	localIndex := index
	parentIndex := localIndex/2
	for parentIndex>0 && parentIndex!=localIndex{
		//fmt.Printf("parent: %d, child: %d\n", parentIndex, localIndex)
		current := q.arr[localIndex]
		parent := q.arr[parentIndex]

		if parent.priority <= current.priority {
			q.unlockAll2(index)
			return
		}
		tmp := parent.priority
		parent.priority = current.priority
		current.priority = tmp
		localIndex = parentIndex
		parentIndex = localIndex/2
	}
	q.unlockAll2(index)
	//fmt.Printf("insert finished: %d\n", priority)
}

func (q *PriorityQueue)RemoveMin()int{
	return 0
}

func (q *PriorityQueue)Dump(name string){
	fmt.Printf("dump starting\n")
	d := lib.NewDot(name)
	for i:=1; 2*i<= q.last; i++{
		//fmt.Printf("dumping: %d\n", i)
		parent := q.arr[i]
		left := q.arr[2*i]
		d.AppendEdge(parent, left)
		if 2*i+1<=q.last{
			right := q.arr[2*i+1]
			d.AppendEdge(parent, right)
		}
	}
	d.Layout()
	fmt.Printf("dump finished\n")
}
