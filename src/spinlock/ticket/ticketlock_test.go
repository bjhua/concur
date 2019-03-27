package ticket

import (
	"lib"
	"testing"
)

func BenchmarkTicketLock1(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 1, sl)
	}
}

func BenchmarkTicketLock2(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 2, sl)
	}
}

func BenchmarkTicketLock3(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 3, sl)
	}
}

func BenchmarkTicketLock4(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 4, sl)
	}
}

func BenchmarkTicketLock5(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 5, sl)
	}
}

func BenchmarkTicketLock6(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 6, sl)
	}
}


func BenchmarkTicketLock7(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 7, sl)
	}
}


func BenchmarkTicketLock8(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 8, sl)
	}
}


func BenchmarkTicketLock16(t *testing.B) {
	sl := NewTicketLock()
	t.N = 1

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 16, sl)
	}
}


func BenchmarkTicketLock32(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 32, sl)
	}
}


func BenchmarkTicketLock48(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 48, sl)
	}
}


func BenchmarkTicketLock64(t *testing.B) {
	sl := NewTicketLock()

	for i:=0; i<t.N; i++{
		lib.BenchLockN(t, 64, sl)
	}
}



