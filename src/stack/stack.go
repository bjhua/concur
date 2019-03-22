package stack

type Stack interface {
	Push(int)
	Pop()(int, err error)
}

