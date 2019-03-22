package locklessskiplist

import (
	"concur/atomic"
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

////////////////
// basic testing

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
	n1.succs[0] = atomic.NewMarkablePtr(unsafe.Pointer(n3))
	n1.succs[1] = atomic.NewMarkablePtr(unsafe.Pointer(n5))
	n1.succs[2] = atomic.NewMarkablePtr(unsafe.Pointer(n5))
	n3.succs[0] = atomic.NewMarkablePtr(unsafe.Pointer(n5))
	n5.succs[0] = atomic.NewMarkablePtr(unsafe.Pointer(n7))
	n5.succs[1] = atomic.NewMarkablePtr(unsafe.Pointer(n7))
	n5.succs[2] = atomic.NewMarkablePtr(unsafe.Pointer(n9))
	n7.succs[0] = atomic.NewMarkablePtr(unsafe.Pointer(n9))
	n7.succs[1] = atomic.NewMarkablePtr(unsafe.Pointer(n9))
	l := n1
	l.dumpList()
	x := 5
	fmt.Printf("lookup(%d) with full path\n", x)
	preds, succs, found := l.lookup(x)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("found=%t\n", found)
	x = 7
	fmt.Printf("lookup(%d) with preds and full path\n", x)
	preds, succs, found = l.lookup(x)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("found=%t\n", found)
	x = 4
	fmt.Printf("lookup(%d) with preds full path\n", x)
	preds, succs, found = l.lookup(4)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("found=%t\n", found)
	x = 6
	fmt.Printf("lookup(%d) with preds and full path\n", x)
	preds, succs, found = l.lookup(6)
	dumpSlice("preds", preds)
	dumpSlice("succs", succs)
	fmt.Printf("found=%t\n", found)
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
	if h.numItems!=10{
		t.Errorf("bug, want: %d, but got: %d\n", 10, h.numItems)
	}
	h.Dump()

}

func TestLazySkipList_SequentialInsert100(t *testing.T) {
	const N = 100

	h := NewSkipList()

	fmt.Printf("insert 0~9\n")
	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	if h.numItems!=N{
		t.Errorf("bug, want: %d, but got: %d\n", N, h.numItems)
	}
	h.Dump()
	found := h.Lookup(51)
	if !found{
		t.Errorf("bug")
	}
}

func TestSkipList_ConcurrentInsert100(t *testing.T) {
	h := NewSkipList()
	//h.Dump()

	const N= 100
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
	if h.numItems != N {
		panic("error")
	}
	fmt.Printf("concurrent insertion:\n")
	h.Dump()
}

func TestSkipList_ConcurrentInsert1000(t *testing.T) {
	h := NewSkipList()
	//h.Dump()

	const N= 1000
	wg := &sync.WaitGroup{}
	wg.Add(2 * N)
	fmt.Printf("insert 0~1000\n")
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
	if h.numItems != N {
		panic("error")
	}
	fmt.Printf("concurrent insertion:\n")
	h.Dump()
	for i:=0; i<N; i++{
		found := h.Lookup(i)
		if !found{
			t.Errorf("bug")
		}
	}
}


func TestSkipList_ConcurrentInsert(t *testing.T) {
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
	for i:=0; i<N; i++{
		found := h.Lookup(i)
		if !found{
			t.Errorf("bug")
		}
	}
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
		b := h.Delete(i)
		if !b{
			panic("error")
		}
	}
	h.Dump()
	if h.numItems!=0{
		panic("error")
	}
}

func TestLazySkipList_SequentialInsertDeleteInsert(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	// sequentially

	const N = 10
	fmt.Printf("insert 0~%d\n", N)
	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	h.Dump()

	fmt.Printf("delete 0~%d\n", N)
	for i := 0; i < N; i++ {
		b := h.Delete(i)
		if !b{
			panic("error")
		}
	}
	h.Dump()
	if h.numItems!=0{
		panic("error")
	}
	fmt.Printf("insert 0~%d\n", N)
	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	h.Dump()
}


func TestLazySkipList_ConcurrentDelete(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	// sequentially

	fmt.Printf("insert 0~9\n")
	const N = 10

	for i := 0; i < N; i++ {
		h.Insert(i)
	}
	h.Dump()

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Delete(j)
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

func TestLazySkipList_ConcurrentInsertDelete(t *testing.T) {
	h := NewSkipList()
	h.Dump()
	fmt.Printf("insert 0~9\n")
	const N = 5

	wg := &sync.WaitGroup{}
	wg.Add(2*N)
	for i := 0; i < N; i++ {
		go func(j int) {
			h.Delete(j)
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

func TestSkipList_Compact(t *testing.T) {
	h := NewSkipListSweep()
	for i:=0; i<1000; i++{
		h.Insert(i)
	}
	h.Dump()
	for i:=0; i<1000; i++{
		h.Delete(i)
	}
	h.Dump()
	time.Sleep(1*time.Millisecond)
	h.Dump()
	time.Sleep(1*time.Millisecond)
	h.Dump()

}

/*

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
				h.Delete(j)
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

*/






