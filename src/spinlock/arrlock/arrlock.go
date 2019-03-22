package arrlock

import (
	"sync/atomic"
)

const maxlen = 1024

type ArrLock struct{
	next int64
	current int64
	array [maxlen]bool
}

func New()*ArrLock{
	l := &ArrLock{}
	l.next = -1
	l.array[0] = true
	return l
}

func (lock *ArrLock)Lock() {
	next := atomic.AddInt64(&lock.next, 1) % maxlen

	for !lock.array[next] {
	}
	lock.current = next
}

func (lock *ArrLock)Unlock(){
	lock.array[lock.current] = false
	lock.array[(lock.current + 1)%maxlen] = true
}









