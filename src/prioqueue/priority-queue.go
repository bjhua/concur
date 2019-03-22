package prioqueue

type PriorityQueue interface {
	Insert(key int)
	RemoveMin()(int, error)
}

type Item interface {
	Equals(interface{})bool
}

// a sample integer matching
type MyInt int

func (i MyInt)Equals(j interface{})bool{
	switch v := j.(type) {
	case int:{
		return (int)(i)==v
	}
	default:
		panic("bad type")
	}
}

