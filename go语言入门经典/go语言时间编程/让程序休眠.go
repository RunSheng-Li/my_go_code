package main

import (
	"fmt"
	"time"
)

// 休眠3s
func main() {
	time.Sleep(3 * time.Second)
	fmt.Println("我睡醒了")
}
