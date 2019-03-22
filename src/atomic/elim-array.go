package atomic

type ElimArray struct{
	arr []*Exchanger
}

func NewElimArray(caps int)*ElimArray{
	l :=  &ElimArray{arr: make([]*Exchanger, 0, caps)}
	for i:=0; i<caps; i++{
		l.arr = append(l.arr, NewExchanger())
	}
	return l
}

func (r *ElimArray)Exchange(index int, role Role, value int)(int, error) {
	var timeOut int64 = 10
	return r.arr[index].Exchange(role, value, timeOut)
}

func (r *ElimArray)Len()int{
	return len(r.arr)
}

