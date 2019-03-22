package golang

import (
	"sync"
	"testing"
)

func Benchmark_RWlockRlock(b *testing.B){
	b.StopTimer()
	lock := &sync.RWMutex{}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		lock.RLock()
		lock.RUnlock()
	}
}

func Benchmark_RWlockRlockMany(b *testing.B){
	b.StopTimer()
	lock := &sync.RWMutex{}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		for j:=0; j<100000000; j++{
			lock.RLock()
		}
		//lock.RUnlock()
	}
}

func Benchmark_RWlockWlock(b *testing.B){
	b.StopTimer()
	lock := &sync.RWMutex{}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		lock.Lock()
		lock.Unlock()
	}
}

func Benchmark_Lock(b *testing.B){
	b.StopTimer()
	lock := &sync.Mutex{}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		lock.Lock()
		lock.Unlock()
	}
}

func Benchmark_RWlockMany(b *testing.B){
	const N = 100000000
	b.StopTimer()
	var locks [N]*sync.RWMutex
	for i:=0; i<N; i++{
		locks[i] = &sync.RWMutex{}
	}

	b.StartTimer()
	for i:=0; i<b.N; i++{
		for j:=0; j<N; j++{
			locks[j].Lock()
		}
		for j:=0; j<N; j++{
			locks[j].Unlock()
		}
	}
}


func Benchmark_RWlockAdd(b *testing.B){
	b.StopTimer()


	b.StartTimer()
	for i:=0; i<b.N; i++{
		_ = 3+4
	}
}