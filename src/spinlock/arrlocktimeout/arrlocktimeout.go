package arrlocktimeout

import (
	"fmt"
	"sync/atomic"
	"time"
)

const arrmaxlen = 20000

type TimeOutLocker interface {
	TryLock(timeout int64)bool
}

//type lockStaus int32

const(
	FREE int32 = iota
	LOCKED
	TIMEOUT
)

type ArrLock struct{
	next int64
	current int64
	array [arrmaxlen]int32
}

func New()*ArrLock{
	l := &ArrLock{}
	l.next = -1
	l.current = -1
	l.array[0] = LOCKED
	return l
}

func (lock *ArrLock)Lock() {
	next := atomic.AddInt64(&lock.next, 1) % arrmaxlen

	//fmt.Printf("in lock()\n")
	//dump(lock)
	for lock.array[next] != LOCKED{
		// spin
	}
	lock.current = next
}


func (lock *ArrLock)TryLock(timeout int64)bool{
	next := atomic.AddInt64(&lock.next, 1) % arrmaxlen
	start := time.Now().UnixNano()
	for{
		var now = time.Now().UnixNano()
		if now - start >= timeout{
			var swapped = atomic.CompareAndSwapInt32(&lock.array[next], FREE, TIMEOUT)
			if swapped{
				return false
			}
	//		fmt.Printf("in trylock()\n")
	//		dump(lock)
			//lock.current = -1
			break
		}
		if lock.array[next] == LOCKED{
			break
		}
	}

	lock.current = next
	return true
}

func output(n int64){
	fmt.Printf("current=%d\n", n)
	time.Sleep(time.Duration(1)*1000)

}

func (lock *ArrLock)Unlock(){
	//output(lock.current)
	lock.array[lock.current] = FREE
	cur := lock.current
	lock.current = -1
	for {
		cur++
		var swapped = atomic.CompareAndSwapInt32(&lock.array[cur % arrmaxlen], FREE, LOCKED)
		if swapped{
			return
		}
		lock.array[cur % arrmaxlen] = FREE
	}
}

//////////
func dump(this *ArrLock){
	var i int64
	for i =0; i<this.next; i++{
		fmt.Printf("[%d]=%d, ", i, this.array[i])
	}
	fmt.Printf("\n")
}



