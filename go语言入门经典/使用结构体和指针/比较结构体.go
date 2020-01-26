//对结构体进行比较，要先看它们的类型和值是否相同
//对于类型相同的结构体，可使用相等性运算符来比较
//要判断两个结构体是否相等，可使用==；要判断它们是否不等，可使用!=
//不能对两个类型不同的结构体进行比较，否则将导致编译错误
//因此，试图比较两个结构体之前，必须确定它们的类型相同。可使用go语言包的reflect
package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {

	a := Drink{
		Name: "可乐",
		Ice:  true,
	}

	b := Drink{
		Name: "可乐",
		Ice:  true,
	}

	if a == b {
		fmt.Println("这是两个相同的结构体")
	}

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
