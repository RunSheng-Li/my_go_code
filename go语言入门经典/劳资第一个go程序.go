package main

import (
	"fmt"
)

func main() {
	fmt.Println("倚天照海花无数，流水高山心自知")

}

//go的主要目标是兼具python等动态语言的开发速度和C或C++等编译语言的性能和安全性
//go可用来创建性能卓越的web服务器和系统程序
//go是编译型语言，要创建并运行go程序，必须执行如下步骤：1.使用文本编辑器创建go程序 2.保存文件 3.编译程序 4.运行编译得到的可执行文件
//go build 启动编译,把我们的包和相关的依赖编译成一个可执行的文件
//go run xxx.go 不同于go build,go run 不会创建可执行文件。go run 提供了一种便利的方式，因为没有必要将编译和执行步骤分开
//go gopher是go语言的吉祥物