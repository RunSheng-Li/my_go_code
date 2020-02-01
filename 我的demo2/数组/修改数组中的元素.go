package main

import "fmt"

func main() {

	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 根据下标访问元素和修改元素内容
	a1[6] = 666

	fmt.Println("a1", a1)

}
