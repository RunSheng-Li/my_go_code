package main

import "fmt"

func main() {

	// 开辟10长度的整型数组
	var a1 [10]int
	fmt.Println("a1", a1)

	// 开辟10长度的整型数组
	var a2 = [10]int{}
	fmt.Println("a2", a2)

	// 创建10长度的整型数组，并给前6项赋值
	var a3 = [10]int{0, 1, 2, 3, 4, 5}
	fmt.Println("a3", a3)

	// 根据实际元素的个数创建数组，其长度依然是固定不变的
	var a4 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a4", a4)

	// 同a4,只是声明方式不一样
	a5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a5", a5)

}
