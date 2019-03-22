package golang

import (
	"fmt"
	"sync"
	"testing"
)

func TestRecursiveLock(t *testing.T){
	lock := &sync.Mutex{}

	lock.Lock()
	fmt.Print("1st\n")
	//lock.Lock()
	fmt.Print("2nd\n")


	//lock.Unlock()
	fmt.Print("1st end\n")
	lock.Unlock()
	fmt.Print("2nd end\n")
}