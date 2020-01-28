package main

import "fmt"

func sumNumbers(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number

	}
	return total

}

func main() {

	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("结果是%v\n", result)

}

//不定参数函数是参数数量不确定的函数，这意味着它们接受可变数量的参数
//在go语言中，能够传递可变数量的参数，但它们的类型必须与函数签名指定的类型相同
//要指定不定参数，可使用三个点...