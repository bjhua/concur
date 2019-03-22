package bounded_queue

type Queue struct{
	ch chan int
}

func New(cap int)*Queue{
	return &Queue{ch:make(chan int, cap)}
}

func (q *Queue)Enq(x int){
	q.ch <- x
}

func (q *Queue)Deq()int{
	n := <-q.ch
	return n
}
