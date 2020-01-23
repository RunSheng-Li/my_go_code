//所谓强类型语言，指的是错误使用了类型时，编译器将引发错误
//所谓动态类型语言，指的是为了执行程序，运行时会将一种类型转换为另一种类型，或者编译器没有实现类型系统

// 类型简介
//package main
//
//import (
//	"fmt"
//)
//
//func sayHello(s string) string {
//	return "Hello " + s
//}
//
//func main() {
//	fmt.Println(sayHello("Sam"))
//}

//go语言中的类型
package main

import (
	"fmt"
)

func addition(x int, y int) int {
	return x + y

}

func main() {
	fmt.Println(addition(1, 4))

	//传递类型不正确的参数 下面这段代码会报错 错误原因是因为在需求int的地方使用了字符串
	//var s string ="fuck"
	//fmt.Println(addition(1, s))

}
