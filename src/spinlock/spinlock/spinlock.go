package spinlock

import (
	"sync/atomic"
)

type SpinLock struct{
	locked int32
}

func NewSpinLock()*SpinLock{
	return &SpinLock{}
}

func (lock *SpinLock)Lock(){
	for atomic.SwapInt32(&(lock.locked), 1) == 1{
	}
}

func (lock *SpinLock)Unlock(){
	lock.locked = 0
}

