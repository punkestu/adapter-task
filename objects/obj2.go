package objects

type Obj2 struct {
	data string
}

func NewObj2(data string) Obj2 {
	return Obj2{data}
}

func (o Obj2) GetData() string {
	return o.data
}

type Obj1Adapter struct {
	obj Obj2
}

func NewObj1Adapter(obj Obj2) Obj1Adapter {
	return Obj1Adapter{obj}
}

func (oa Obj1Adapter) GetData() int {
	return len(oa.obj.GetData())
}
