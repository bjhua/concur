package test

import (
	"concur/spinlock"
	"fmt"
	"sync"
	"testing"
)

func Test_MCSLock2(t *testing.T){
	var i = 0
	lock := spinlock.NewMcsLock2()
	//lock := &sync.Mutex{}
	var loops = 10000
	wg := &sync.WaitGroup{}

	f := func(){
		defer wg.Done()

		for k:=0; k<loops; k++{
			lock.Lock()
			i++
			lock.Unlock()
		}
	}
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	if i != 2*loops{
		fmt.Printf("i=%d\n", i)
		t.Fatal("wrong")
	}
}
