package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	var s string = "egg"
	intToString := strconv.Itoa(i)
	var breakfast string = intToString + s
	fmt.Println(breakfast)

}

//只能拼接类型为字符串的变量，如果整数和字符串进行拼接将导致编译错误
//go标准库提供了strconv包，您可使用其中的方式Itoa来完成这种任务：它将整数转为字符串