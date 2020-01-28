package main

import "fmt"

func main() {

	defer fmt.Println("我在函数结束后运行")
	fmt.Println("heheda")

}

//使用一条defer语句，在它所在的函数执行完毕后执行另一个函数
