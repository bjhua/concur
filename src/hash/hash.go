package hash

type HashSet interface {
	Insert(key int)bool
	Lookup(key int)bool
	Delete(key int)bool
	Items()int
}

type HashItem interface {
	HashCode()int
	Equals(interface{})bool
}

// a sample integer matching
type MyInt int
func (i MyInt)HashCode()int{
	return (int)(i)
}

func (i MyInt)Equals(j interface{})bool{
	switch v := j.(type) {
	case int:{
		return (int)(i)==v
	}
	default:
		panic("bad type")
	}
}

