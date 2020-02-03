package main

import "fmt"

//定义结构体
type Dog struct {
	//定义狗的属性
	name string
	sex  int
}

//封装狗的方法
func (d *Dog) SetName(name string) {
	//在这里可以加一些权限校验
	d.name = name
}

func (d *Dog) GetName() string {
	//在这里可以加一些权限校验
	return d.name
}

//定义咬人方法
func (d *Dog) bite() {

	fmt.Printf("让本汪%s来给你一些教育。。。\n", d.name)

}

func main() {

	//创建空白对象给属性赋值
	var d = Dog{}
	//给对象赋值
	d.name = "旺财"
	d.sex = 1
	fmt.Println(d.sex)
	d.bite()

	//创建对象并有选择地给属性赋值
	var d2 = Dog{name: "哮天犬"}
	d2.sex = 1
	d2.bite()

	//创建空白指针对象并给属性赋值
	var dp = new(Dog)
	dp.SetName("土狗")
	dp.bite()

}
