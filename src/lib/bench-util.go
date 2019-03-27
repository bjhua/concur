package lib

import (
	"sync"
	"testing"
)

const limit = 10000

func BenchLockN(b *testing.B, numThreads int, lock sync.Locker)int64{
	wg := &sync.WaitGroup{}
	wg.Add(numThreads)
	var local int64 = 0

	for i:=0; i<numThreads; i++{
		go func(){
			for j:=0; j<limit; j++{
				lock.Lock()
				//local++
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	if local != (int64)(numThreads*limit){
		//panic("wrong")
	}
	return local
}

