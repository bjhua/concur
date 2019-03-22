package atomic

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Status int
type Role int

type Exchanger struct{
	lock *sync.Mutex

	status Status
	role Role
	inValueCh chan int
	outValueCh chan int
	outCh chan interface{}
}

const (
	Status_Empty Status = iota
	Status_Waiting
	Status_Paired
)

const(
	Role_None Role = iota
	Role_Yin
	Role_Yang
)

func NewExchanger()*Exchanger{
	return &Exchanger{lock: &sync.Mutex{},
		role: Role_None,
		status: Status_Empty,
		inValueCh: make(chan int, 1),
		outValueCh: make(chan int, 1),
		outCh: make(chan interface{}, 1),
	}
}

func (r *Exchanger)Exchange(role Role, value int, nanoTimeOut int64)(int, error) {
	startTime := time.Now().UnixNano()

	RETRY:
	r.lock.Lock()
	switch r.status {
	case Status_Empty:
		{
			fmt.Printf("who %d, with value %d enter empty\n", role, value)

			r.role = role
			r.status = Status_Waiting
			r.inValueCh <- value
			r.lock.Unlock()

			select {
			case <-r.outCh:
				{
					//fmt.Printf("who %d, with value %d got value\n", role, value)
					v := <-r.outValueCh
					r.role = Role_None
					r.status = Status_Empty
					return v, nil
				}
			case <-func() chan bool {
				ch := make(chan bool, 1)
				time.Sleep(time.Duration(nanoTimeOut) * time.Nanosecond)
				ch <- true
				return ch
			}():
				{
					//fmt.Printf("who %d, with value %d timeout\n", role, value)

					r.lock.Lock()
					if r.status == Status_Paired {
						_ = <-r.outCh
						v := <-r.outValueCh
						r.status = Status_Empty
						r.role = Role_None
						r.lock.Unlock()
						return v, nil
					}
					_ = <-r.inValueCh
					r.status = Status_Empty
					r.role = Role_None
					r.lock.Unlock()
					return 0, errors.New("timeout: no partner arrived")
				}
			}
		}
	case Status_Waiting:
		{
			if r.role == role {
				r.lock.Unlock()
				now := time.Now().UnixNano()
				if now-startTime>nanoTimeOut{
					return 0, errors.New("timeout: same kind of op")
				}
				goto RETRY
			}
			r.status = Status_Paired
			n := <-r.inValueCh
			r.outValueCh <- value
			r.outCh <- struct {
			}{}
			r.lock.Unlock()
			return n, nil
		}
	case Status_Paired:
		{
			r.lock.Unlock()
			nowTime := time.Now().UnixNano()
			if nowTime-startTime>nanoTimeOut{
				return 0, errors.New("paired")
			}
			goto RETRY
		}
	default:
		{
			r.lock.Unlock()
			panic("impossible")
		}
	}
}


