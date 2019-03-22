package test

import (
	"sync"
	"testing"
)

const limit = 10000

func benchLockN(b *testing.B, n int, lock sync.Locker)int64{
	wg := &sync.WaitGroup{}

	var local int64 = 0

	for i:=0; i<n; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()

			for {
				lock.Lock()
				
				//fmt.Printf("local = %d [%d]\n", local, limitx)
				if local >= limit{
					lock.Unlock()
					return
				}else{
					local++
					lock.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	if local != limit{
		panic("wrong")
	}
	//fmt.Printf("vvvv = %d\n", local)
	return local
}

