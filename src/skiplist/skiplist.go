package skiplist

type SkipList interface {
	Insert(key int)bool
	Lookup(key int)bool
	Delete(key int)bool
	Items()int
}

type SkipListItem interface {
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

