package elim_stack2

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

const N = 100

func TestStack(t *testing.T) {
	wg := &sync.WaitGroup{}
	q := New()
	wg.Add(N)
	for i:=0; i<N; i++{
		go func(j int){
			q.Push(j)
			wg.Done()
		}(i*i)
	}

	wg.Wait()
	fmt.Printf("push finished")

	sl := make([]int, 0, 10)
	for i:=0; i<N; i++{
		var n = 0
		var err error = nil

		for{
			n, err = q.Pop()
			if err==nil{
				goto L
			}
		}
	L:
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
