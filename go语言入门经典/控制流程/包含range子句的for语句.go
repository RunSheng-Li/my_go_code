package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println("循环到", i)
		fmt.Println(n)
	}

}

//声明变量numbers,并将一个包含4个整数的数组赋给它
//for语句指定了迭代变量i，用于存储索引值。这个变量将在每次迭代结束后更新
//for语句指定了迭代变量n，用于存储来自数组中的值。它也将在每次迭代结束后更新
//在循环中，打印这两个变量的值
