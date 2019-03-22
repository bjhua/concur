package spinlock

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type TtasBackoff struct{
	locked int64
	bo []string
}

func NewTtasBackoff()*TtasBackoff{
	rand.Seed(time.Now().Unix())
	v := &TtasBackoff{}
	return v
}

func backoffGen(min int32, max int32)func(){
	var limit = min
	return func() {
		delay := rand.Int31n(limit)
		if max > 2*limit{
			limit = 2*limit
		}else{
			limit = max
		}
		time.Sleep(time.Duration(delay)*time.Microsecond)
	}
}

func (lock *TtasBackoff)Lock(){
	backOff := backoffGen(1, 512)
	for{
		for atomic.LoadInt64(&lock.locked) == 1{
		}

		if atomic.SwapInt64(&lock.locked, 1) == 0{
			return
		}
		backOff()
	}

	//fmt.Printf("goroutine [%d] lock held: %s\n", lock.rid, lock.name)
}

func (lock *TtasBackoff)Unlock(){
	lock.locked = 0
	//fmt.Printf("goroutine [%d] lock released: %s\n", tid, lock.name)
}









