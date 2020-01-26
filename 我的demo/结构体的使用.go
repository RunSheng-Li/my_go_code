package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func main() {

	//顺序初始化，必须全部初始化完整
	var man Person = Person{"sam", 'g', 18}
	fmt.Println(man)

	//部分初始化
	man2 := Person{sex: 'm', age: 19}
	fmt.Println(man2)

	//索引成员变量 “.”
	fmt.Printf("%q\n", man2.name)

	var man3 Person
	man3.name = "sam"
	man3.sex = 'a'
	man3.age = 99
	fmt.Println(man3)
	man3.age = 100
	fmt.Println(man3)


}

/*

结构体：
	是一种数据  类型。

	type Person struct {		—— 类型定义 （地位等价于 int byte bool string ....） 通常放在全局位置。
		name string
		sex  byte
		age int
	}

普通变量定义和初始化：

	1. 顺序初始化: 依次将结构体内部所欲成员初始化。

		var man Person = Person{"andy"， 'm',  20}

	2. 指定成员初始化：

		man := Person{name:"rose", age:18}		---- 未初始化的成员变量，取该数据类型对应的默认值

普通变量的赋值和使用：

	使用“.”索引成员变量。

		var man3 Person
		man3.name = "mike"
		man3.sex = 'm'
		man3.age = 99

*/
