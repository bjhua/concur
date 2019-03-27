package spinlock

import (
	"fmt"
	"lib"
	"testing"
)

func BenchmarkSpinLock1(b *testing.B) {
	const numThreads = 32
	sl := NewSpinLock()

	for thd:=1; thd<numThreads; thd++{
		b.Run(fmt.Sprintf("numThread-%d", thd), func (b *testing.B){
			for i:=0; i<b.N; i++{
				lib.BenchLockN(b, thd, sl)
			}
		})
	}
}


/*
func BenchmarkSpinLock2(t *testing.B) {
	sl := NewSpinLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 2, sl)
	}
}

func BenchmarkSpinLock3(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 3, sl)
	}
}

func BenchmarkSpinLock4(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 4, sl)
	}
}

func BenchmarkSpinLock5(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 5, sl)
	}
}

func BenchmarkSpinLock6(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 6, sl)
	}
}


func BenchmarkSpinLock7(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 7, sl)
	}
}


func BenchmarkSpinLock8(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 8, sl)
	}
}


func BenchmarkSpinLock16(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 16, sl)
	}
}


func BenchmarkSpinLock32(t *testing.B) {
	sl := NewSpinLock()
	//t.N = 1000

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 32, sl)
	}
}


func BenchmarkSpinLock48(t *testing.B) {
	sl := NewSpinLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 48, sl)
	}
}


func BenchmarkSpinLock64(t *testing.B) {
	sl := NewSpinLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 64, sl)
	}
}


func BenchmarkSpinLock128(t *testing.B) {
	sl := NewSpinLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 128, sl)
	}
}


func BenchmarkSpinLock256(t *testing.B) {
	sl := NewSpinLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 256, sl)
	}
}


*/


