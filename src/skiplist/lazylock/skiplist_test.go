package lazyskiplist

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

////////////////
// basic testing
func TestList_condense(t *testing.T){
	n1 := NewNode(1, 2)
	n2 := NewNode(2, 1)
	n3 := NewNode(3, 0)

	// no dup
	ns := make([]*list, 0, 1)
	ns = append(ns, n1)
	ns = append(ns, n2)
	ns = append(ns, n3)
	nr := condense(ns, 3)
	if len(nr)!=3||
		nr[0].data != 3 ||
		nr[1].data !=2 ||
		nr[2].data !=1{
		panic("error")
	}
	// dup
	ns = make([]*list, 0, 1)
	ns = append(ns, n1)
	ns = append(ns, n1)
	ns = append(ns, n2)
	ns = append(ns, n2)
	ns = append(ns, n3)
	ns = append(ns, n3)
	nr = condense(ns, 6)
	if len(nr)!=3||
		nr[0].data != 3 ||
		nr[1].data != 2 ||
		nr[2].data != 1{
		panic("error")
	}
}

func TestList_lookup(t *testing.T) {
	/*
	1       5         9
	1       5    7    9
	1   3   5    7    9
	 */
	n1 := NewNode(1, 3)
	n3 := NewNode(3, 1)
	n5 := NewNode(5, 3)
	n7 := NewNode(7, 2)
	n9 := NewNode(9, 3)
	n1.succs[0] = n3
	n1.succs[1] = n5
	n1.succs[2] = n5
	n3.succs[0] = n5
	n5.succs[0] = n7
	n5.succs[1] = n7
	n5.succs[2] = n9
	n7.succs[0] = n9
	n7.succs[1] = n9
	l := n1
	x := 5
	fmt.Printf("lookup(%d) with full path\n", x)
	preds, succs, topLevel := l.lookup(x,true)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("topLevel=%d\n", topLevel)
	x = 7
	fmt.Printf("lookup(%d) with preds and full path\n", x)
	preds, succs, topLevel = l.lookup(x, true)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("topLevel=%d\n", topLevel)
	x = 7
	fmt.Printf("lookup(%d) without full path\n", x)
	preds, succs, topLevel = l.lookup(7, false)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("topLevel=%d\n", topLevel)
	x = 4
	fmt.Printf("lookup(%d) without full path\n", x)
	preds, succs, topLevel = l.lookup(4, false)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("topLevel=%d\n", topLevel)
	x = 6
	fmt.Printf("lookup(%d) with full path\n", x)
	preds, succs, topLevel = l.lookup(6, true)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("topLevel=%d\n", topLevel)
	//os.Exit(1)
}

func TestLazySkipList_SequentialInsert(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	// sequentially

	x := 8
	fmt.Printf("insert %d\n", x)
	h.Insert(x)
	h.Dump()

	x = 9
	fmt.Printf("insert %d\n", x)
	h.Insert(x)
	h.Dump()

	x = 9
	fmt.Printf("insert %d\n", x)
	h.Insert(x)
	h.Dump()

	fmt.Printf("insert 0~9\n")
	for i := 0; i < 10; i++ {
		h.Insert(i)
	}
	h.Dump()

}

func TestSkipList_ConcurrrentInsert(t *testing.T) {
	h := NewSkipList()
	//h.Dump()

	const N = 10000
	wg := &sync.WaitGroup{}
	wg.Add(2 * N)
	fmt.Printf("insert 0~100\n")
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	if h.numItems != N{
		panic("error")
	}
	fmt.Printf("concurrent insertion:\n")
	h.Dump()
}

func TestSkipList_SeqLookup(t *testing.T) {
	h := NewSkipList()

	const N = 1000
	wg := &sync.WaitGroup{}
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("concurrent insertion finished:\n")
	h.Dump()
	b := h.Lookup(50)
	if !b{
		panic("error")
	}
	b = h.Lookup(1000)
	if b{
		panic("error")
	}
}


func TestSkipList_ConLookup(t *testing.T) {
	h := NewSkipList()

	const N = 100
	wg := &sync.WaitGroup{}
	wg.Add(2 * N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("concurrent insertion:\n")
	h.Dump()

	wg = &sync.WaitGroup{}
	wg.Add(2*N)
	for i := 0; i < N; i++ {
		go func(j int) {
			b := h.Lookup(j)
			if !b {
				panic("error")
			}
			wg.Done()
		}(i)
		go func(j int) {
			b := h.Lookup(N+1+j)
			if b{
				panic("error")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("concurrent insertion:\n")
	h.Dump()
}


func TestLazySkipList_SequentialDelete(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	// sequentially

	const N = 10
	fmt.Printf("insert 0~9\n")
	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	h.Dump()

	for i := 0; i < N; i++ {
		b := h.Delete(0, i)
		if !b{
			panic("error")
		}
	}
	h.Dump()
	if h.numItems!=0{
		panic("error")
	}
}

func TestLazySkipList_ConcurrentDelete(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	// sequentially

	fmt.Printf("insert 0~9\n")
	const N = 4

	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	h.Dump()

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Delete(j, j)
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Printf("concurrent delete:\n")
	h.Dump()
	if h.numItems!=0{
		t.Errorf("num: %d, want %d\n", h.numItems, 0)
	}
}

func TestLazySkipList_ConInsertDelete(t *testing.T) {
	h := NewSkipList()
	fmt.Printf("insert 0~9\n")
	const N = 100

	wg := &sync.WaitGroup{}
	wg.Add(2*N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Delete(j, j)
			wg.Done()
		}(i)
		go func(j int) {
			h.Insert(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("concurrent insert and delete:\n")
	h.Dump()

}

//////////////////////////
// benchmark

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


func BenchmarkSkipList_ConcurrentLookup(b *testing.B) {
	const max = 1000000 //1000,0000

	b.StopTimer()
	h := NewSkipList()
	for i := 0; i < max; i++ {
		n := (int)(rand.Int31())
		h.Insert(n)
	}
	fmt.Printf("create %d\n", max)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bench(h)
	}
}

const maxItems = 2000000

func BenchmarkSkipList_ConInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewSkipList()
		wg := &sync.WaitGroup{}
		wg.Add(maxItems)
		for i := 0; i < maxItems; i++ {
			go func(j int){
				h.Insert(j)
				wg.Done()
			}(i)
		}
		wg.Wait()
		if h.numItems != maxItems{
			b.Errorf("want: %d, but got: %d\n", maxItems, h.numItems)
		}
	}
}

func BenchmarkSkipList_ConDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewSkipList()
		wg := &sync.WaitGroup{}
		wg.Add(maxItems)
		for i := 0; i < maxItems; i++ {
			go func(j int){
				h.Insert(j)
				wg.Done()
			}(i)
		}
		wg.Wait()
		if h.numItems != maxItems{
			b.Errorf("want: %d, but got: %d\n", maxItems, h.numItems)
		}

		wg = &sync.WaitGroup{}
		wg.Add(maxItems)
		for i := 0; i < maxItems; i++ {
			go func(j int){
				h.Delete(0, j)
				wg.Done()
			}(i)
		}
		wg.Wait()
		if h.numItems != 0{
			b.Errorf("want: %d, but got: %d\n", maxItems, h.numItems)
		}
	}
}

func BenchmarkSkipList_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewSkipList()
		for i := 0; i < maxItems; i++ {
			n := (int)(rand.Int31())
			h.Insert(n)
		}
	}
}
