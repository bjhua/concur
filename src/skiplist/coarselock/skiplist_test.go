package coarseskiplist

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestLockHash(t *testing.T) {
	h := NewSkipList()
	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()

	// sequentially
	h.Insert(88)
	h.Insert(99)
	h.Insert(72)

	fmt.Printf("items = %d\n", h.numItems)
	h.Dump()

	h.Delete(99)
	fmt.Printf("items after delete 99 = %d\n", h.numItems)
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

func bench(h *SkipList) {
	const numRoutines = 10000
	const numOpsPerRoutine = 10000

	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func(j int) {
			for j:=0; j<numOpsPerRoutine; j++ {
				n := (int)(rand.Int31())

				if j%10 == 0 {
					//_ = h.Insert(n)
				} else {
					_ = h.Lookup(n)
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}



func BenchmarkNewSkipListenchmarkSkipList(b *testing.B) {
	const max = 1000000 //1000,0000

	b.StopTimer()
	h := NewSkipList()
	for i:=0; i<max; i++{
		n := (int)(rand.Int31())
		h.Insert(n)
	}
	fmt.Printf("create %d\n", max)


	b.StartTimer()
	for i:=0; i<b.N; i++{
		bench(h)
	}
}

const maxItems = 10000000
func BnchmarkSkipList_Insert(b *testing.B) {
	for i:=0; i<b.N; i++{
		h := NewSkipList()
		for i:=0; i<maxItems; i++{
			n := (int)(rand.Int31())
			h.Insert(n)
		}
	}
}
