//在go语言中，声明变量时如果没有给它指定值，则变量将为默认值，这种默认值被称为零值
//变量的默认值取决于其类型

//变量的零值
//------------------------------------------------------------------------------
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var i int
//	var f float64
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, f, b, s)
//}

//检查变量的值是否为零值
//------------------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	var s string
	if s == "" {
		fmt.Printf("s 是零值")
	}
}

//在go语言中，为确定变量是否已经赋值，不能检查它是否为nil，而必须检查它是否为默认值
//由于类型string的零值为"",因此对于类型为string的变量，要确定是否已经给它赋值，可检查其值是否为零值""
