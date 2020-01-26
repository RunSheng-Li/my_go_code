package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a

	a = 100
	fmt.Println("a = ", a)

	*p = 250 // 借助a变量的地址，操作a对应的空间
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)

	a = 1000
	fmt.Println("*p = ", *p)

}

//指针就是地址。指针变量就是存储地址的变量
//*p 解引用 间接引用
//操作符“&”取变量地址，“*”通过指针访问目标对象
//等号 左边的变量，代表 变量所指向的内存空间。	（写）
// 等号 右边的变量，代表 变量内存空间存储的数据值。	（读）