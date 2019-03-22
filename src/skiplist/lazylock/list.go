package lazyskiplist

import (
	"fmt"
	"math"
	"sync"
)

type list struct {
	*sync.Mutex

	data      int
	succs     []*list
	deleted   bool
	numLevels int //
}

func NewNode(x int, maxLevel int) *list {
	n := &list{Mutex: &sync.Mutex{}, data: x, numLevels: maxLevel}
	n.succs = make([]*list, maxLevel, maxLevel)
	return n
}

func NewList(maxCaps int) *list {
	tail := NewNode(math.MaxInt32, maxCaps)
	head := NewNode(math.MinInt32, maxCaps)
	for i := 0; i < maxCaps; i++ {
		head.succs[i] = tail
	}
	return head
}

// methods
// Lookup some item in the given skiplist
// return false if not found
// return true if found with its predessor and successor
// Example:
/*
For the following skip:

      1,       5,       9
      1,    4, 5,       9
      1, 2, 4, 5        9
      1, 2, 4, 5, 6, 7, 9

lookup(5) will return:
[1, 1, 2, 2], [5, 5, 5, 5], true
 */
func (l *list) dumpList() {
	n := l
	for n != nil {
		fmt.Printf("%d[%d], ", n.data, len(n.succs))
		n = n.succs[0]
	}
	fmt.Printf("\n")
}

func dumpSlice(s string, sl []*list) {
	fmt.Printf("%s[", s)
	for _, n := range sl {
		fmt.Printf("%d, ", n.data)
	}
	fmt.Printf("]\n")
}

func appendNode(preds []*list, succs []*list, p *list, s *list) ([]*list, []*list) {
	return append(preds, p), append(succs, s)
}

func (l *list) lookup(x int, withFullPath bool) ([]*list, []*list, int) {
	var preds, succs []*list = nil, nil
	var cur *list = nil
	topLevel := -1

	if withFullPath {
		preds = make([]*list, 0, 1)
		succs = make([]*list, 0, 1)
	}
	pred := l
	for level := l.numLevels - 1; level >= 0; level-- {
		cur = pred.succs[level]

		for cur.data < x {
			pred = cur
			cur = cur.succs[level]
		}
		if withFullPath {
			preds, succs = appendNode(preds, succs, pred, cur)
		}
		if cur.data == x {
			if topLevel == -1 {
				topLevel = level
			}
			if !withFullPath {
				return preds, succs, topLevel
			}
		}
	}
	return preds, succs, topLevel
}

func condense(ns []*list, maxLevel int) []*list {
	r := make([]*list, 0, 1)
	nsLen := len(ns)
	r = append(r, ns[nsLen-1])
	last := 0
	for i := 1; i < maxLevel; i++ {
		if r[last] == ns[nsLen-1-i] {
			continue
		}
		r = append(r, ns[nsLen-1-i])
		last++
	}
	return r
}

func validate(preds []*list, succs []*list, numLevels int) bool {
	predsLen := len(preds)
	for i := 0; i < numLevels; i++ {
		p := preds[predsLen - 1 - i]
		s := succs[predsLen - 1 - i]

		if p.deleted || s.deleted || p.succs[i] != s {
			return false
		}
	}
	return true
}

func validate2(preds []*list, target *list, numLevels int) bool {
	predsLen := len(preds)
	for i := 0; i < numLevels; i++ {
		p := preds[predsLen-1-i]

		if p.deleted || target.deleted || p.succs[i] != target {
			return false
		}
	}
	return true
}

func (l *list) insert(x int, numLevels int) bool {
RETRY:
	preds, succs, topLevel := l.lookup(x, true)
	if topLevel != -1 {
		return false
	}
	nonDupPreds := condense(preds, numLevels)
	for _, p := range nonDupPreds {
		p.Lock()
	}
	if !validate(preds, succs, numLevels) {
		for _, p := range nonDupPreds {
			p.Unlock()
		}
		goto RETRY
	}

	fresh := NewNode(x, numLevels)
	predLen := len(preds)
	for i := 0; i < numLevels; i++ {
		fresh.succs[i] = succs[predLen-1-i]
	}
	for i := 0; i < numLevels; i++ {
		preds[predLen-1-i].succs[i] = fresh
	}
	for _, p := range nonDupPreds {
		p.Unlock()
	}
	return true
}

func (l *list) delete(tid int, x int) bool {
RETRY:
	preds, succs, topLevel := l.lookup(x, true)
	if topLevel == -1 {
		return false
	}
	if topLevel+1 != succs[len(succs)-1-topLevel].numLevels {
		goto RETRY
	}
	nonDupPreds := condense(preds, topLevel+1)
	target := succs[len(succs)-1-topLevel]
	target.Lock()
	for _, p := range nonDupPreds {
		p.Lock()
	}

	if !validate2(preds, target, topLevel+1) {
		for _, p := range nonDupPreds {
			p.Unlock()
		}
		target.Unlock()
		goto RETRY
	}

	predLen := len(preds)
	for i := 0; i <= topLevel; i++ {
		preds[predLen-1-i].succs[i] = target.succs[i]
	}
	target.deleted = true
	for _, p := range nonDupPreds {
		p.Unlock()
	}
	target.Unlock()
	return true
}
