package coarseskiplist

import (
	"math"
)

type data struct{
	data int
}

type node struct {
	*data
	next *node
	down *node
	pred *node
}

/*    --------
      | next |--->
  <---| pred |
      | down |
      ---|----
         |
         \/
 */
type list struct {
	node *node
	//tail *node
}

func NewList() *list {
	tail := &node{data: &data{data: math.MaxInt32}}
	head := &node{data: &data{data: math.MinInt32}, next: tail}
	tail.pred = head
	return &list{node: head}
}

// methods
// Lookup some item in the given skiplist
// return false if not found
// return true if found with the full path
// Example:
/*
For the following skip:

      1,       5,       9
      1,    4, 5,       9
      1, 2, 4, 5        9
      1, 2, 4, 5, 6, 7, 9

lookupWithPath(4) will return:
[1, 1, 2, 2], true
 */
func (l *list) lookupWithPath(x int) ([]*node, bool) {
	path := make([]*node, 0, 1)

	pred := l.node
	cur := pred.next

L:
	for cur.data.data < x {
		pred = cur
		cur = cur.next
	}
	if cur.data.data > x {
		if pred.down == nil {
			path = append(path, pred)
			return path, false
		}
		path = append(path, pred)
		pred = pred.down
		cur = pred.next
		goto L
	}
	// cur.data == x
	path = append(path, pred)
	cur = cur.down
	for cur != nil{
		path = append(path, cur.pred)
		cur = cur.down
	}
	return path, true
}

func (l *list) lookupRaw(x int) bool {
	pred := l.node
	cur := pred.next
	for {
		for cur.data.data < x {
			pred = cur
			cur = cur.next
		}
		if cur.data.data == x {
			return true
		}
		if pred.down == nil {
			return false
		}

		pred = pred.down
		cur = pred.next
	}
}

func (l *list) insert(x int, maxLevel int) bool {
	path, found := l.lookupWithPath(x)
	if found {
		return false
	}
	var pred, n, down *node = nil, nil, nil

	v := &data{data: x}
	for i := len(path) - 1; i >= 0; i-- {
		pred = path[i]
		n = &node{data: v, next: pred.next, down: down, pred: pred}
		pred.next.pred = n
		pred.next = n
		down = n
		maxLevel--
		if maxLevel == 0 {
			break
		}
	}
	return true
}

func (l *list) delete(x int) bool {
	path, found := l.lookupWithPath(x)
	if !found {
		return false
	}
	var pred *node = nil
	for i := len(path) - 1; i >= 0; i-- {
		pred = path[i]
		if pred.next.data.data != x {
			break
		}
		cur := pred.next
		pred.next = cur.next
		cur.next.pred = pred
		cur.next, cur.pred = nil, nil
	}
	return true
}
