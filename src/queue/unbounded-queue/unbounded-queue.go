package unbounded_queue

import "errors"

type Queue struct{
	ch chan []int
}

func New()*Queue{
	q := &Queue{ch:make(chan []int, 1)}
	q.ch <- make([]int, 0, 10)
	return q
}

func (q *Queue)Enq(x int){
	sl := <-q.ch
	sl = append(sl, x)
	q.ch <- sl
}

func (q *Queue)Deq()(int, error){
	sl := <-q.ch
	if len(sl)==0{
		return 0, errors.New("empty")
	}
	n := sl[0]
	q.ch <- sl[1:]
	return n, nil
}
