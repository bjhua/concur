package lib

// the channel based implementation of semaphore
type Semaphore struct{
	ch chan struct{}
}

func NewSemaphore(limit int)*Semaphore{
	sem := &Semaphore{ch: make(chan struct{}, limit)}
	return sem
}

func (sem *Semaphore)Acquire(){
	sem.ch <- struct {}{}
}

func (sem *Semaphore)Release(){
	<-sem.ch
}

