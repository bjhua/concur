package test

import (
	"spinlock"
	"sync"
	"testing"
)

//const limi = 1000000

func benchTasLockN(t *testing.B, n int, lock *spinlock.Ttas,
	limitx int64)int64{
	wg := &sync.WaitGroup{}


	var local int64 = 0

	for i:=0; i<n; i++{
		wg.Add(1)
		go func(){
			for {
				lock.Lock(2)
				//fmt.Printf("local = %d [%d]\n", local, limitx)
				if local >= limitx{
					lock.Unlock()
					goto L
				}else{
					local = local+1
					lock.Unlock()
				}
			}
		L:
			wg.Done()
		}()
	}

	wg.Wait()
	//fmt.Printf("vvvv = %d\n", local)
	return local
}

func BenchmarkTasLock1(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	//acc = 0
	t.N = 1000

	_ = benchTasLockN(t, 1, sl, limit)
	//fmt.Printf("value = %d\n\n\n", vv)
}

func BenchmarkTasLock2(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 2, sl, limit)
}

func BenchmarkTasLock3(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 3, sl, limit)
}

func BenchmarkTasLock4(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 4, sl, limit)
}

func BenchmarkTasLock5(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 5, sl, limit)
}

func BenchmarkTasLock6(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 6, sl, limit)
}


func BenchmarkTasLock7(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 7, sl, limit)
}


func BenchmarkTasLock8(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 8, sl, limit)
}


func BenchmarkTasLock16(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 16, sl, limit)
}


func BenchmarkTasLock32(t *testing.B) {
	sl := spinlock.NewTtas("spinlock")
	t.N = 1000

	benchTasLockN(t, 32, sl, limit)
}

/*
func BenchmarkSpinLock48(t *testing.B) {
	sl := spinlock.New("spinlock")

	benchSpinLockN(t, 48, sl)
}


func BenchmarkSpinLock64(t *testing.B) {
	sl := spinlock.New("spinlock")

	benchSpinLockN(t, 64, sl)
}
*/


