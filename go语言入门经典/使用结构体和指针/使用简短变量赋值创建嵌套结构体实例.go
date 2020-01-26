//有时候，数据结构需要包含多个层级
//为了建立较复杂的数据结构，在一个结构体中嵌套另一个结构体的方式很有用
//例如：超级英雄列表：对于每位超级英雄，都需要存储其住址，而住址本身也是一个数据结构，非常适合使用结构体来表示

package main

import "fmt"

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

func main() {

	e := Superhero{
		Name: "钢铁侠",
		Age:  18,
		Address: Address{
			Number: 1008,
			Street: "深圳市福田区",
			City:   "深圳",
		},
	}
	fmt.Printf("%+v\n", e)
	fmt.Println(e.Address.Street)

}
