package main

import "fmt"

//在切片中添加元素 在切片末尾添加多个元素
func main() {

	var haha = make([] string, 2)
	haha[0] = "你好"
	haha[1] = "sam"
	haha = append(haha, "添加的元素")
	fmt.Println(haha)
	fmt.Println(haha[2])
	haha = append(haha,"1","2","3")
	fmt.Println(haha)

}
