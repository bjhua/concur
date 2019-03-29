package spinlock

import (
	"lib"
	"sync/atomic"
)

type TtasBackoff struct {
	locked int64
}

func NewTtasBackoff() *TtasBackoff {
	v := &TtasBackoff{}
	return v
}

func (lock *TtasBackoff) Lock() {
	backOff := lib.BackoffGen(1, 512)
	for {
		for atomic.LoadInt64(&lock.locked) == 1 {
		}

		if atomic.SwapInt64(&lock.locked, 1) == 0 {
			return
		}
		backOff()
	}

	//fmt.Printf("goroutine [%d] lock held: %s\n", lock.rid, lock.name)
}

func (lock *TtasBackoff) Unlock() {
	lock.locked = 0
	//fmt.Printf("goroutine [%d] lock released: %s\n", tid, lock.name)
}
