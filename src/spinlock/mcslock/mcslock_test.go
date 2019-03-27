package mcslock

import (
	"fmt"
	"lib"
	"sync"
	"testing"
)

func Test_MCSLock(t *testing.T){
	var i = 0
	lock := New()
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

func localBench(locker sync.Locker){
	fmt.Printf("locker start===\n")
	locker.Lock()
	locker.Unlock()
	fmt.Printf("locker finished===\n")
}

func BenchmarkNewMcsLock1(t *testing.B) {
	lock := New()
	//lock := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 1, lock)
	}
	//fmt.Printf("value = %d\n\n\n", vv)
}

/*
func BenchmarkMcsLock2(t *testing.B) {
	sl := spinlock.NewMcsLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 2, sl)
	}
}

func BenchmarkMcsLock3(t *testing.B) {
	sl := spinlock.NewMcsLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 3, sl)
	}
}

func BenchmarkMcsLock4(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 4, sl)
	}
}

func BenchmarkMcsLock5(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 5, sl)
	}
}

func BenchmarkMcsLock6(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 6, sl)
	}
}


func BenchmarkMcsLock7(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 7, sl)
	}
}


func BenchmarkMcsLock8(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 8, sl)
	}
}


func BenchmarkMcsLock16(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 16, sl)
	}
}


func BenchmarkMcsLock32(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 32, sl)
	}
}

func BenchmarkMcsLock48(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 48, sl)
	}
}


func BenchmarkMcsLock64(t *testing.B) {
	sl := spinlock.NewMcsLock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 64, sl)
	}
}

*/

