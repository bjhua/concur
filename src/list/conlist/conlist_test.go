package conlist

import "testing"

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