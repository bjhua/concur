package ticket

import (
	"sync/atomic"
)

type TicketLock struct{
	next int64
	current int64
}

func NewTicketLock()*TicketLock{
	return &TicketLock{next: 0, current: 1}
}

func (lock *TicketLock)Lock(){
	n := atomic.AddInt64(&lock.next, 1)
	for !atomic.CompareAndSwapInt64(&lock.current, n, n) {
	}

	//fmt.Printf("goroutine [%d] lock held: %s\n", lock.rid, lock.name)
}

func (lock *TicketLock)Unlock(){
	lock.current++
	//tid := lock.rid
	//fmt.Printf("goroutine [%d] lock released: %s\n", tid, lock.name)
}









