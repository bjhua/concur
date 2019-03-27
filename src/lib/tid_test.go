package lib

import (
	"fmt"
	"sync"
	"testing"
)

func TestTidGen(t *testing.T){
	const numRoutines = 100
	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)
	for i:=0; i<numRoutines; i++{
		go func(){
			tid := TidGen()
			fmt.Printf("%d\n", tid)
			wg.Done()
		}()
	}
	wg.Wait()

}


func TestTidEqual(t *testing.T){
	tid := TidGen()
	if tid == TidNone{
		t.Errorf("bad tid\n")
	}
}