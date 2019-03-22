package queue

type Queue interface {
	Enq(x int)
	Deq()int
}


