package main

import "fmt"

func main(){

	//声明实际变量
	var a int = 20
	//声明指针变量
	var ip *int

	//指针变量的存储地址
	ip = &a

	fmt.Printf("a 变量的地址是: %x\n", &a  )
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )
	fmt.Printf("*ip 变量的值: %d\n", *ip )


}