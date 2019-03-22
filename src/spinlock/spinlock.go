package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct{
	locked int32
	LLL int
}

func NewSpinLock()*SpinLock{
	runtime.GOMAXPROCS(30)

	return &SpinLock{}
}

func (lock *SpinLock)Lock(){
	for atomic.SwapInt32(&(lock.locked), 1) == 1{
	}
}

func (lock *SpinLock)Unlock(){
	lock.locked = 0
}

