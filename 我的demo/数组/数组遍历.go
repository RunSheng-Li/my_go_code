package main

import "fmt"

func main() {

	a1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 按照下标访问元素
	fmt.Println(a1[6])
	fmt.Printf("type=%T,len=%d\n", a1, len(a1))

	// 通过下标遍历数组元素
	for i := 0; i < len(a1); i++ {
		fmt.Println("下标=", i, "值=", a1[i])
	}

	// 通过range遍历数组元素
	///index为遍历中的元素下标，value为值
	for index, value := range a1 {
		fmt.Println(index, value)
	}

}
