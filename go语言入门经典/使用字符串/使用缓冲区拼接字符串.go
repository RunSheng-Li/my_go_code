package main

import (
	"bytes"
	"fmt"
)

func main() {

	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")

	}

	fmt.Println(buffer.String())

}

//创建一个空的字节缓冲区，并将其赋值给变量buffer
//一个运行500次的循环，每次循环都将“z”写入缓冲区
//循环结束后，对缓冲区调用函数String()以字符串的方式输出结果