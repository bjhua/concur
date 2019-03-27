package test

import (
	"spinlock"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var rid int32 = 0
func genRid()int32{
	n := atomic.AddInt32(&rid, 1)
	fmt.Printf("n=%d\n", n)
	return n
}

var wg = &sync.WaitGroup{}
var counter int32 = 0
var rsplock *spinlock.RSpinLock = spinlock.NewRSpinLock()


func f(rid int32){
	var acc = 0
	defer wg.Done()

	for j:=0; j<10000; j++{
		for k:=0; k<10; k++{
			rsplock.Lock(rid)
		}
		acc++
		for k:=0; k<10; k++ {
			rsplock.Unlock()
		}
	}
}

func Test_RSpinLock(t *testing.T) {
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go f(genRid())
	go f(genRid())
	wg.Wait()
	fmt.Printf("i=%d\n", counter)
}
