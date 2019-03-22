package coarselock

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestLockHash(t *testing.T) {
	h := New()
	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()
	// sequentially
	h.Insert(88)
	h.Insert(99)
	h.Insert(72)

	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()

	h.Delete(99)
	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()

	h.Delete(72)
	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()


	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i:=0; i<100; i++{
		go func(i int) {
			h.Insert(104 + i*16)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("items = %d ==? 103\n", h.numItems)
	h.Dump()

	wg = &sync.WaitGroup{}
	wg.Add(100)
	for i:=0; i<100; i++{
		go func(i int) {
			h.Delete(104 + i*16)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("items = %d ==? 103\n", h.numItems)
	h.Dump()

}

const numOps = 10000000

func bench(h *HashSet){
	wg := &sync.WaitGroup{}
	wg.Add(numOps)

	for i:=0; i<numOps; i++{
			go func(j int){
				n := rand.Int()
				if j%10==0{
					_ = h.Insert(n)
				}else{
					_ = h.Lookup(n)
				}

				wg.Done()
			}(i)
	}

	wg.Wait()
}

const max = 10000000 //1000,0000

func BenchmarkHashSet(b *testing.B) {
	b.StopTimer()
	h := New()
	for i:=0; i<max; i++{
		n := rand.Int()
		h.Insert(n)
	}
	//h.Dump()


	b.StartTimer()
	for i:=0; i<b.N; i++{
		bench(h)
	}
}
