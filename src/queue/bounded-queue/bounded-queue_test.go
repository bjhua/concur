package bounded_queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New(10)
	q.Enq(1)
	q.Enq(2)
	q.Enq(3)

	n1, n2, n3 := q.Deq(), q.Deq(), q.Deq()
	if n1!=1 || n2 !=2 || n3 !=3{
		panic("failed")
	}
}
