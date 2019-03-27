package coarselockprioqueue

import (
	"sync"
	"testing"
)

func checkIt(t *testing.T, q *PriorityQueue, rootIndex int){
	root := q.arr[rootIndex]
	leftIndex := rootIndex*2
	rightIndex := rootIndex*2 + 1
	if leftIndex<=q.last{
		leftChild := q.arr[leftIndex]
		if root.priority > leftChild.priority{
			t.Errorf("bug")
		}
		checkIt(t, q, leftIndex)
	}
	if rightIndex<=q.last{
		rightChild := q.arr[rightIndex]
		if root.priority > rightChild.priority{
			t.Errorf("bug")
		}
		checkIt(t, q, rightIndex)
	}
}

func check(t *testing.T, q *PriorityQueue){
	checkIt(t, q, 1)
}

func TestPioQueue(t *testing.T) {
	const N = 10
	q := NewPrioQueue(N)
	for i:=0; i<N; i++{
		q.Insert(20+i)
	}
	check(t, q)
	q.Dump("test_1")

}

func Test_ConInsert10(t *testing.T) {
	const N = 10
	q := NewPrioQueue(N)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i:=0; i<N; i++{
		go func(j int){
			q.Insert(1+j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	q.Dump("test_con_insert_10")

}

func Test_ConInsert100(t *testing.T) {
	const N = 100
	q := NewPrioQueue(N)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i:=0; i<N; i++{
		go func(j int){
			q.Insert(1+j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	q.Dump("test_con_insert_100")

}

func Test_ConInsert100000(t *testing.T) {
	const N = 100000
	q := NewPrioQueue(N)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i:=0; i<N; i++{
		go func(j int){
			q.Insert(1+j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}

func Test_ConInsert1000000_single(t *testing.T) {
	const N = 1000000
	q := NewPrioQueue(N)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i:=0; i<N; i++{
		go func(j int){
			q.Insert(1+j)
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}


func Test_ConInsert1000000(t *testing.T) {
	const N = 100
	const numRoutines = 10000
	q := NewPrioQueue(N*numRoutines)
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(j int){
			for k:=0; k<N; k++{
				q.Insert(j*N + k)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}

func Test_ConInsert1000000_2(t *testing.T) {
	const N = 100000
	const numRoutines = 10
	q := NewPrioQueue(N*numRoutines)
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(j int){
			for k:=0; k<N; k++{
				q.Insert(j*N + k)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}


func Test_ConInsert10000000(t *testing.T) {
	const N = 1000
	const numRoutines = 10000
	q := NewPrioQueue(N*numRoutines)
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(j int){
			for k:=0; k<N; k++{
				q.Insert(j*N + k)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}

func Test_ConInsert10000000_2(t *testing.T) {
	const N = 1000000
	const numRoutines = 10
	q := NewPrioQueue(N*numRoutines)
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(j int){
			for k:=0; k<N; k++{
				q.Insert(j*N + k)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	check(t, q)
	//q.Dump("test_con_insert_100000")

}


func Test_ConInsert100000000_2(t *testing.T) {
	const N = 10000000
	const numRoutines = 10
	q := NewPrioQueue(N*numRoutines)
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(j int){
			for k:=0; k<N; k++{
				q.Insert(j*N + k)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	//check(t, q)
	//q.Dump("test_con_insert_100000")

}




/*
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
*/

