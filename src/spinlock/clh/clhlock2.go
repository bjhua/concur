package clh

import (
	"sync/atomic"
	"unsafe"
)

type node2 struct{
	released bool
	prev *node2
}

type McsLock2 struct{
	tail *AtomicRef  // tail of the queue
	current *node2 // current holder of the lock
}

func NewMcsLock2()*McsLock2{
	l := &McsLock2{tail: NewAtomicRef(unsafe.Pointer(&node2{released:true}))}
	return l
}

func (lock *McsLock2)Lock() {
	fresh := &node2{released:false}
	var oldTail = lock.tail.getAndSet(unsafe.Pointer(fresh))
	for !(*node2)(oldTail).released{
		//spin
	}
	oldTail = nil // to make GC happy
	lock.current = fresh
}

func (lock *McsLock2)Unlock() {
	lock.current.released = true
}

type AtomicRef struct{
	ref unsafe.Pointer
}

func NewAtomicRef(ref unsafe.Pointer)*AtomicRef{
	return &AtomicRef{ref: ref}
}

func (this *AtomicRef)getAndSet(p unsafe.Pointer)unsafe.Pointer{
	return atomic.SwapPointer(&this.ref, p)
}







