package objects

type Obj1 struct {
	data int
}

func NewObj1(data int) Obj1 {
	return Obj1{data}
}

func (o Obj1) GetData() int {
	return o.data
}
