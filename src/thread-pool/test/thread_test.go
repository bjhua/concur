package test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestThread(t *testing.T)  {
	var max int64 = 10000000
	var count int64 = 0
	var i int64
	for i=0; i<max; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			time.Sleep(1000 *time.Second)
			//wg.Done()
		}()
	}
	fmt.Printf("%d\n", count)
	for count!=max{
	}
	fmt.Printf("%d\n", count)

}

func BenchmarkThread(b *testing.B){
	wg := &sync.WaitGroup{}

	for i:=0; i<b.N; i++{
		wg.Add(1)
		go func(){
			wg.Done()
		}()
		wg.Wait()
	}
}