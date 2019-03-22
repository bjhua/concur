package atomic

import (
	libatomic "sync/atomic"
	"unsafe"
)

type pair struct{
	marked bool
	ptr unsafe.Pointer
}

type MarkablePtr struct{
	pair unsafe.Pointer // is always *pair
}

func New(ptr unsafe.Pointer)*MarkablePtr{
	return &MarkablePtr{pair: unsafe.Pointer(&pair{marked:false, ptr:ptr})}
}

func (r *MarkablePtr)Get()(bool, unsafe.Pointer){
	pair := (*pair)(r.pair)
	return pair.marked, pair.ptr
}

func (r *MarkablePtr)CompareAndSwap(oldPtr, newPtr unsafe.Pointer,
	oldMarked, newMarked bool)bool{
		current := (*pair)(r.pair)
	if current.marked != oldMarked || current.ptr != oldPtr{
		return false
	}
	if current.marked == newMarked && current.ptr == newPtr{
		return true
	}
	newPair := &pair{marked: newMarked, ptr: newPtr}
	return libatomic.CompareAndSwapPointer(&r.pair, unsafe.Pointer(current), unsafe.Pointer(newPair))
}




