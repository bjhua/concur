package thread_pool

type Stack interface {
	Push(int)
	Pop()(int, err error)
}

