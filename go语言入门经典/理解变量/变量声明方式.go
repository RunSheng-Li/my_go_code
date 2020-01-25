//变量声明方式,下面列出了所有的变量声明方式
//------------------------------------------------------------------------------

//var s string = "hello"
//var s1 = "hello"
//var t string
//t = "hello"
//u := "hello"
//该用哪种方式呢？go对此有一定的限制——不能在函数外面使用简短变量声明。在遵守这条规则的前提下，怎么做都可以

package main

import (
	"fmt"
)

var s = "hello"

func main() {
	i := 66
	fmt.Println(s)
	fmt.Println(i)
}
