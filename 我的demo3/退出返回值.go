package main

import (
	"fmt"
	"os"
)

func main(){

	fmt.Println("go语言中的main函数不支持任何返回值")
	fmt.Println("通过os.Exit来返回状态")
	os.Exit(1)


}