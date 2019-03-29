package test

import (
	"fmt"
	"sync"
	"testing"
	time2 "time"
)

func TestRoutine(t *testing.T){
	const numPerRoutine = 100
	const numRoutine = 1000000
	var arr [numPerRoutine*numRoutine]int

	wg := &sync.WaitGroup{}
	wg.Add(numRoutine)
	for r:=0; r<numRoutine; r++{
		go func(id int){
			sum := 0
			for i:=0; i<numPerRoutine; i++{
				sum += arr[numPerRoutine * id + i]
			}
			wg.Done()
		}(r)
	}
	wg.Wait()

}

func TestRoutineRun(t *testing.T){
	wg := &sync.WaitGroup{}
	wg.Add(1)
	start := time2.Now().Nanosecond()
	go func() {
		wg.Done()
	}()
	wg.Wait()
	end := time2.Now().Nanosecond()
	fmt.Printf("%d", end-start)
}
