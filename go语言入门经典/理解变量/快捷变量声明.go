//快捷变量声明
//------------------------------------------------------------------------------
//package main
//
//import (
//	"fmt"
//)
//
//func main(){
//	var s,t string = "foo","bar"
//	fmt.Println(s)
//	fmt.Println(t)
//
//}

//以快捷方式声明类型不同的变量
//------------------------------------------------------------------------------
package main

import (
	"fmt"
)

func main() {
	var (
		s string = "foo"
		i int    = 6
	)
	fmt.Println(s)
	fmt.Println(i)

}

// 声明变量后，就不能再次声明它。虽然可以给变量重新赋值，但不能重新声明变量，否则将导致编译阶段错误
