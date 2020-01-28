package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for {
			fmt.Println("我是go程哈哈")
		}
	}()

	for {
		runtime.Gosched() //出让当前时间片
		fmt.Println("我是主go程")

	}

}
