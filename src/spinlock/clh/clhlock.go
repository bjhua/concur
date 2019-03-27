package clh

import (
	"sync/atomic"
	"unsafe"
)

type node struct{
	holdingOrWaiting bool
	prev *node
}

type McsLock struct{
	tail unsafe.Pointer  // tail of the queue
	current *node // current holder of the lock
}

func NewMcsLock()*McsLock{
	l := &McsLock{tail: unsafe.Pointer(&node{holdingOrWaiting:false, prev:nil})}
	return l
}

func (lock *McsLock)Lock() {
	fresh := &node{holdingOrWaiting:true, prev:nil}
	p := unsafe.Pointer(fresh)
	var oldTail unsafe.Pointer = atomic.SwapPointer(&lock.tail, p)
	for (*node)(oldTail).holdingOrWaiting{
		//spin
	}
	oldTail = nil // to make GC happy
	lock.current = fresh
}

func (lock *McsLock)Unlock() {
	lock.current.holdingOrWaiting = false
}





