package test

import (
	"sync"
	"testing"
)

//const limi = 1000000

func benchMutLockN(t *testing.B, n int, lock sync.Locker,
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
	//fmt.Printf("vvvv = %d\n", local)
	return local
}

func BenchmarkMutLock1(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 1, sl, limit)
	}
	//fmt.Printf("value = %d\n\n\n", vv)
}

func BenchmarkMutLock2(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 2, sl, limit)
	}
}

func BenchmarkMutLock3(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 3, sl, limit)
	}
}

func BenchmarkMutLock4(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 4, sl, limit)
	}
}

func BenchmarkMutLock5(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 5, sl, limit)
	}
}

func BenchmarkMutLock6(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 6, sl, limit)
	}
}


func BenchmarkMutLock7(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 7, sl, limit)
	}
}


func BenchmarkMutLock8(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 8, sl, limit)
	}
}


func BenchmarkMutLock16(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 16, sl, limit)
	}
}


func BenchmarkMutLock32(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 32, sl, limit)
	}
}

func BenchmarkMutLock48(t *testing.B) {
	sl := &sync.Mutex{}

	for i:=0; i<t.N; i++{
		benchMutLockN(t, 48, sl, limit)
	}
}


func BenchmarkMutLock64(t *testing.B) {
	sl := &sync.Mutex{}

	benchMutLockN(t, 64, sl, limit)
}



