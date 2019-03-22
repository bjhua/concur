package lockless_queue

import (
	"fmt"
	"sort"
	"testing"
)

const N = 10000

func TestQueue(t *testing.T) {
	q := New()
	for i:=0; i<N; i++{
		go func(j int){
			q.Enq(j)
		}(i*i)
	}

	sl := make([]int, 0, 10)
	for i:=0; i<N; i++{
		n, _ := q.Deq()
		sl = append(sl, n)
		fmt.Printf("%d, ", n)
	}

	if len(sl)!=N{
		panic("wrong")
	}

	fmt.Printf("\n\n")

	sort.Ints(sl)
	for _, n := range sl{
		fmt.Printf("%d, ", n)
	}
}
