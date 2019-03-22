package spinlock

import (
	"sync/atomic"
)

type Ttas struct{
	value int64
	rid int64
	name string
}

func NewTtas(name string)*Ttas{
	return &Ttas{value: 0, name: name}
}

func (lock *Ttas)Lock(rid int64){
	for{
		for atomic.LoadInt64(&lock.value) == 1{
		}
		if atomic.SwapInt64(&lock.value, 1) == 0{
			break
		}
	}

	lock.rid = rid
	//fmt.Printf("goroutine [%d] lock held: %s\n", lock.rid, lock.name)
}

func (lock *Ttas)Unlock(){
	lock.rid = 0
	lock.value = 0
	//tid := lock.rid
	//fmt.Printf("goroutine [%d] lock released: %s\n", tid, lock.name)
}









