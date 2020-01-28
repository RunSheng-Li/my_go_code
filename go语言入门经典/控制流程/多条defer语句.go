package main

import "fmt"

func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("heheda")
}

//3条defer语句都指定了它们所在的函数执行完毕后要执行的函数
//外部函数执行完毕后，按与defer语句出现顺序相反的顺序执行它们指定的函数
