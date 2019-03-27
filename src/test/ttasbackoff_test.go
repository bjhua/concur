package test

import (
	"spinlock"
	"sync"
	"testing"
)

//const limi = 1000000

func benchTasBackoffLockN(t *testing.B, n int, lock *spinlock.TtasBackoff,
	limitx int64)int64{
	wg := &sync.WaitGroup{}


	var local int64 = 0

	for i:=0; i<n; i++{
		wg.Add(1)
		go func(){
			for {
				lock.Lock()
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
	if local != limit{
		panic("wrong")
	}
	//fmt.Printf("vvvv = %d\n", local)
	return local
}

func BenchmarkTasBackoffLock1(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	_ = benchTasBackoffLockN(t, 1, sl, limit)
	//fmt.Printf("value = %d\n\n\n", vv)
}

func BenchmarkTasBackoffLock2(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 2, sl, limit)
}

func BenchmarkTasBackoffLock3(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 3, sl, limit)
}

func BenchmarkTasBackoffLock4(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 4, sl, limit)
}

func BenchmarkTasBackoffLock5(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 5, sl, limit)
}

func BenchmarkTasBackoffLock6(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 6, sl, limit)
}


func BenchmarkTasBackoffLock7(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 7, sl, limit)
}


func BenchmarkTasBackoffLock8(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 8, sl, limit)
}


func BenchmarkTasBackoffLock16(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 16, sl, limit)
}


func BenchmarkTasBackoffLock32(t *testing.B) {
	sl := spinlock.NewTtasBackoff()
	t.N = 1000

	benchTasBackoffLockN(t, 32, sl, limit)
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


