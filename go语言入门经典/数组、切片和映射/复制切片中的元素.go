package main

import "fmt"

func main() {

	var haha = make([] string, 3)
	haha[0] = "1"
	haha[1] = "2"
	var hehe = make([] string, 2)
	copy(hehe, haha)
	fmt.Println(hehe)

}

//函数copy在新切片中创建元素的副本，因此修改一个切片中的元素不会影响另一个切片
