package target

import (
	"fmt"
)

type Obj interface {
	GetData() int
}

type Target struct {
	obj Obj
}

func NewTarget(obj Obj) Target {
	return Target{obj}
}

func (t Target) PrintData() {
	fmt.Printf("Data: %d\n", t.obj.GetData())
}
