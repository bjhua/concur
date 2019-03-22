package test

import (
	"concur/spinlock"
	"fmt"
	"runtime"
	"testing"
)

const maxCores = 28

func BenchmarkSpinLock1(t *testing.B) {
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

func BenchmarkSpinLock2(t *testing.B) {
	sl := spinlock.NewSpinLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 2, sl)
	}
}

func BenchmarkSpinLock3(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 3, sl)
	}
}

func BenchmarkSpinLock4(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 4, sl)
	}
}

func BenchmarkSpinLock5(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 5, sl)
	}
}

func BenchmarkSpinLock6(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 6, sl)
	}
}


func BenchmarkSpinLock7(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 7, sl)
	}
}


func BenchmarkSpinLock8(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 8, sl)
	}
}


func BenchmarkSpinLock16(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 16, sl)
	}
}


func BenchmarkSpinLock32(t *testing.B) {
	sl := spinlock.NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		benchLockN(t, 32, sl)
	}
}


func BenchmarkSpinLock48(t *testing.B) {
	sl := spinlock.NewSpinLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 48, sl)
	}
}


func BenchmarkSpinLock64(t *testing.B) {
	sl := spinlock.NewSpinLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 64, sl)
	}
}


func BenchmarkSpinLock128(t *testing.B) {
	sl := spinlock.NewSpinLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 128, sl)
	}
}


func BenchmarkSpinLock256(t *testing.B) {
	sl := spinlock.NewSpinLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 256, sl)
	}
}





