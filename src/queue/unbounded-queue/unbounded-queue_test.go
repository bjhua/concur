package unbounded_queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	for i:=0; i<100; i++{
		q.Enq(i*i)
	}

	for i:=0; i<100; i++{
		n, _ := q.Deq()
		fmt.Printf("%d, ", n)
	}
}
