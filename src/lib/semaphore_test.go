package lib

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count int32 = 0

func TestSemaphore(t *testing.T) {
	const caps = 3
	sem := NewSemaphore(caps)
	for i:=0; i<caps; i++{
		go func(){
			sem.Acquire()
			atomic.AddInt32(&count, 1)
		}()
	}
	// lazy to use the waitgroup
	time.Sleep(1*time.Second)
	if count!=caps{
		t.Errorf("bad")
	}
	go func(){
		sem.Acquire()
		atomic.AddInt32(&count, 1)
	}()
	time.Sleep(1*time.Second)
	if count!=caps{
		t.Errorf("bad")
	}
	for i:=0; i<caps; i++{
		go func(){
			sem.Release()
			atomic.AddInt32(&count, -1)
		}()
	}
	time.Sleep(1*time.Second)
	if count!=1{
		t.Errorf("bad")
	}
}

func BenchmarkSemaphore(b *testing.B) {
	const numRoutine = 10000
	const numSem = 100
	sem := NewSemaphore(numSem)
	wg := &sync.WaitGroup{}
	b.N = 1
	for i:=0; i<1; i++{
		wg.Add(numRoutine)
		for j:=0; j<numRoutine; j++{
			go func() {
				sem.Acquire()
				time.Sleep(10*time.Microsecond)
				sem.Release()
				wg.Done()
			}()
		}
	}
	wg.Wait()

}


