package main

import "fmt"

func main() {
	var haha = make(map[string]int)
	haha["age"] = 17
	haha["phone"] = 110
	haha["height"] = 180
	fmt.Println(haha["age"])
	fmt.Println(haha)

}

//在其他编程语言中，映射也被称为关联数组、字段或散列
//映射可视为键-值对集合