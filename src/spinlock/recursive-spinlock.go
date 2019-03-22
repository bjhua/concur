package spinlock

type RLocker interface {
	Lock(rid int32)
	Unlock()
}

type RSpinLock struct{
	theLock *SpinLock
	id int32
	count int32
}

func NewRSpinLock()*RSpinLock{
	return &RSpinLock{theLock:NewSpinLock()}
}

func (this *RSpinLock)Lock(rid int32){
	if this.theLock.locked==1 && this.id==rid{
		this.count++
		return
	}
	this.theLock.Lock()
	this.id = rid
	this.count++
}

func (this *RSpinLock)Unlock(){
	this.count--
	if this.count==0{
		this.id = 0
		this.theLock.Unlock()
	}
}

