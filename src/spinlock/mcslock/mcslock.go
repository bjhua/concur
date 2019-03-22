package mcslock

import (
	"sync/atomic"
	"unsafe"
)

type node struct{
	locked bool
	next *node
}

type McsLock struct{
	tail unsafe.Pointer  // tail of the queue
	current unsafe.Pointer // current holder of the lock
}

func New()*McsLock{
	l := &McsLock{}
	return l
}

func (lock *McsLock)Lock() {
	fresh := &node{}
	p := unsafe.Pointer(fresh)
	pred := atomic.SwapPointer(&lock.tail, p)
	if pred == nil{
		lock.current = p
		return
	}
	(*node)(pred).next = fresh
	for !fresh.locked{
	}
	(*node)(pred).next = nil
	lock.current = p
}

func (lock *McsLock)Unlock() {
	if atomic.CompareAndSwapPointer(&lock.tail, lock.current, nil){
		lock.current = nil
		return
	}
	for (*node)(lock.current).next == nil{
	}
	(*node)(lock.current).next.locked = true
}





