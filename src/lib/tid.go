package lib

import (
	"sync/atomic"
)

type Tid int32

const TidNone Tid = -1

var tid Tid = 0

// Generate a unique thread id, starting from 1.
func TidGen() Tid {
	i := atomic.AddInt32((*int32)(&tid), 1)
	return (Tid)(i)
}


