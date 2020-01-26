package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {

	a := Drink{
		Name: "可乐",
		Ice:  true,
	}

	b := &a
	b.Ice = false
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)

}

//要修改原始结构体实例包含的值，必须使用指针
//指针是指向内存地址的引用，因此使用它操作的不是结构体的副本而是其本身
//指针和值的差别很微妙，但选择使用指针还是值很容易区分：如果需要修改原始结构体实例，就使用指针；如果要操作一个结构体，但不想修改原始结构体实例，就使用值
