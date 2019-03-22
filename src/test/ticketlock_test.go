package test

import (
	"concur/spinlock"
	"testing"
)

func BenchmarkTicketLock1(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 1, sl)
	}
}

func BenchmarkTicketLock2(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 2, sl)
	}
}

func BenchmarkTicketLock3(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 3, sl)
	}
}

func BenchmarkTicketLock4(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 4, sl)
	}
}

func BenchmarkTicketLock5(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 5, sl)
	}
}

func BenchmarkTicketLock6(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 6, sl)
	}
}


func BenchmarkTicketLock7(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 7, sl)
	}
}


func BenchmarkTicketLock8(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 8, sl)
	}
}


func BenchmarkTicketLock16(t *testing.B) {
	sl := spinlock.NewTicketLock()
	t.N = 1

	for i:=0; i<t.N; i++{
		benchLockN(t, 16, sl)
	}
}


func BenchmarkTicketLock32(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 32, sl)
	}
}


func BenchmarkTicketLock48(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 48, sl)
	}
}


func BenchmarkTicketLock64(t *testing.B) {
	sl := spinlock.NewTicketLock()

	for i:=0; i<t.N; i++{
		benchLockN(t, 64, sl)
	}
}



