package arrlocktimeout

import (
	"fmt"
	"lib"
	"sync"
	"testing"
)

const loops = 2000

func Test_ArrLockTimeOut(t *testing.T) {
	var acc int64 = 0
	wg := &sync.WaitGroup{}
	lock := New()
	//rcount := make(chan int, 100)

	f := func(n int) {
		defer wg.Done()
		var numTry = 0
		var numLock = 0

		for j := 0; j < 4; j++ {
			var locked = lock.TryLock(10)
			if locked {
				numTry++
				acc++
				lock.Unlock()
			} else {
				lock.Lock()
				numLock++
				acc++
				//fmt.Println("test.unlock()", locked)
				lock.Unlock()
			}
		}
		fmt.Printf("tries: %d, locks: %d\n", numTry, numLock)
	}
	wg.Add(loops)
	for i := 0; i < loops; i++ {
		go f(i)
	}
	wg.Wait()
	if acc != 4*loops {
		fmt.Printf("%d\n", acc)
		panic("wrong")
	}
}

func Benchmark_ArrLockTimeOut(t *testing.B) {
	sl := New()

	for numCore := 1; numCore < 30; numCore++ {
		t.Run(fmt.Sprintf("numCore-%d", numCore), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				lib.BenchLockN(t, numCore, sl)
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
