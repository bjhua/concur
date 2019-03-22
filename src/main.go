package main

import (
	"concur/spinlock"
	"fmt"
	"sync/atomic"
)

var ids int64 = 0

func genId()int64{
	id := atomic.AddInt64(&ids, 1)
	return id
}

func main() {
	sLock := spinlock.NewSpinLock()
	sLock.Lock()

	//lock.Lock()

	sLock.Unlock()

	for i:=0; i<10; i++ {
		go func(rid int64) {
			sLock.Lock()
			fmt.Printf("rid = %d\n", rid)
			sLock.Unlock()
		}(genId())
	}

	for{
		fmt.Print("")
	}



}

