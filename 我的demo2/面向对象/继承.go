package main

import "fmt"

type Doggy struct {
	name string
	sex  bool
}

//定义父类方法
func (d *Doggy) bite() {
	fmt.Printf("%s要咬你了啊\n", d.name)
}

//定义子类
type PoliceDog struct {
	//持有一个Doggy对象，并继承Doggy的全部属性和方法
	Doggy

	//警犬所独有的属性
	skill string
}

//覆写父类方法
func (pd *PoliceDog) bite() {
	fmt.Printf("%s还要咬你就自己挂了\n", pd.name)

}

//发展出新的方法
func (pd *PoliceDog) fuck() {
	fmt.Printf("%sfuck u!!!", pd.name)
}

func main() {

	var d1 = new(PoliceDog)
	d1.skill = "fuck!!!"
	d1.name = "土狗"
	d1.sex = true
	d1.bite()
	d1.fuck()

}
