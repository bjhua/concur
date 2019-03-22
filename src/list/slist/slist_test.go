package slist

import (
	"sync"
	"testing"
	"time"
)

func TestSlist(t *testing.T) {
	l := New()
	l.Dump()
	l.Add(1)
	l.Add(3)
	l.Add(2)
	l.Dump()

	b := l.Find(2)
	if !b{
		panic("wrong")
	}
	b = l.Find(4)
	if b{
		t.Fatal("wrong")
	}
}

func TestSlistConcur(t *testing.T) {
	var loops int32 = 5
	l := New()
	wg := &sync.WaitGroup{}
	wg.Add((int)(loops))
	var i int32
	for i=0; i<loops; i++{
		go func(n int32) {
			defer wg.Done()

			time.Sleep(100*time.Nanosecond)
			l.Add(n)
		}(i)
	}

	wg.Wait()
	l.Dump()
}