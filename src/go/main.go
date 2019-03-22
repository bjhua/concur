package golang

import (
	"fmt"
	"sync"
)

const max = 1000

var rwLock = &sync.RWMutex{}
var n = 0

func main(){
	wg := &sync.WaitGroup{}
	wg.Add(max)

	for i:=0; i<max; i++{
		if i%2==0{
			go func(j int){  // reader
				rwLock.RLock()
				fmt.Printf("thread[%d] rlock\n", j)
				fmt.Printf("n=%d\n", n)
				fmt.Printf("thread[%d] runlock\n", j)
				rwLock.RUnlock()

				wg.Done()
			}(i)
		}else{
			go func(j int){  // writer
				rwLock.Lock()
				fmt.Printf("thread[%d] wlock\n", j)
				n++
				fmt.Printf("thread[%d] wunlock\n", j)
				rwLock.Unlock()

				wg.Done()
			}(i)
		}
	}

	wg.Wait()
	fmt.Printf("n=%d\n", n)
}
