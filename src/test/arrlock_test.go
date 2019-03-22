package test

import (
	"concur/spinlock"
	"concur/spinlock/arrlock"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func Test_ArrLock(t *testing.T){
	i := 0
	wg := &sync.WaitGroup{}
	lock := arrlock.New()

	f := func() {
		defer wg.Done()

		for j:=0; j<10000; j++{
			lock.Lock()
			i++
			lock.Unlock()
		}
	}
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	if i != 2*10000{
		panic("wrong")
	}
}

func BenchmarkAbLock1(t *testing.B) {
	sl := arrlock.New()


	for i:=0; i<t.N; i++{
		benchLockN(t, 1, sl)
	}
	//fmt.Printf("value = %d\n\n\n", vv)
}

func BenchmarkArrLock1(t *testing.B) {
	sl := spinlock.NewSpinLock()

	runtime.GOMAXPROCS(30)

	for numCore:=1; numCore<30; numCore++{
		t.Run(fmt.Sprintf("numCore-%d", numCore), func (b *testing.B){
			for i:=0; i<b.N; i++{
				benchLockN(t, numCore, sl)
			}
		})
	}
}

/*
func BenchmarkAbLock2(t *testing.B) {
	sl := spinlock.NewAblock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 2, sl)
	}
}

func BenchmarkAbLock3(t *testing.B) {
	sl := spinlock.NewAblock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 3, sl)
	}
}

func BenchmarkAbLock4(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 4, sl)
	}
}

func BenchmarkAbLock5(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 5, sl)
	}
}

func BenchmarkAbLock6(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 6, sl)
	}
}


func BenchmarkAbLock7(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 7, sl)
	}
}


func BenchmarkAbLock8(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 8, sl)
	}
}


func BenchmarkAbLockx(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 16, sl)
	}
}


func BenchmarkAbLock32(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 32, sl)
	}
}

func BenchmarkAbLock48(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 48, sl)
	}
}


func BenchmarkAbLock64(t *testing.B) {
	sl := spinlock.NewAblock()
	//t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 64, sl)
	}
}

*/


