package main

import (
	"fmt"

	"github.com/punkestu/adapter-task/objects"
	"github.com/punkestu/adapter-task/target"
)

func main() {
	fmt.Println("+ Percobaan menggunakan struct yang sesuai")
	var obj1 = objects.NewObj1(10)
	var target1 = target.NewTarget(obj1)
	target1.PrintData()

	fmt.Println("\n+ Percobaan menggunakan adapter")
	var obj2 = objects.NewObj2("Hello, World!")
	// var target2 = target.NewTarget(obj2) // Akan error karena tidak sesuai
	var objAdapted = objects.NewObj1Adapter(obj2)
	var target2 = target.NewTarget(objAdapted)
	target2.PrintData()
}
