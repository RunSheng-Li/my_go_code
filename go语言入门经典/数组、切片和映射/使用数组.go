//声明数组并给它赋值
package main

import "fmt"

func main() {
	var haha [2] string
	haha[0] = "hello"
	haha[1] = "sam"
	fmt.Println(haha[0])
	fmt.Println(haha[1])
	fmt.Println(haha)

}

//使用关键字var声明一个为haha的变量
//将一个长度为2的数组赋给这个变量
//这个数组的类型为字符串

//声明数组的长度后，就不能给它添加元素了
