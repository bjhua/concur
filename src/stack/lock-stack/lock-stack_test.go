package elim_stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	sl := make([]int, 0, 1)
	sl = append(sl, 222)

	lent := len(sl)
	fmt.Printf("len=%d, %d\n", lent, sl[lent-1])
	sl = sl[:lent-1]
	lent = len(sl)
	fmt.Printf("len=%d\n", lent)
	for _, v := range sl{
		fmt.Printf("v=%d, ", v)
	}

	k := New()
	k.Push(1)
	k.Push(2)
	k.Push(3)

	n1, err := k.Pop()
	n2, err := k.Pop()
	n3, err := k.Pop()
	if err != nil{
		panic(err.Error())
	}

	if n1!=3 || n2 !=2 || n3 !=1{
		panic("failed")
	}
}
