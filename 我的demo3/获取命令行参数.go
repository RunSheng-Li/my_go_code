package main

import (
	"fmt"
	"os"
)

//在程序中直接通过os.Args获取命令行参数
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello,", os.Args[1])
	}

}
