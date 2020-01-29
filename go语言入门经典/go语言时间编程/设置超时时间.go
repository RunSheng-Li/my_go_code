package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("你有两秒钟计算19*4")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("时间到了！答案是79，你算出来了吗？")
		}
	}
}
