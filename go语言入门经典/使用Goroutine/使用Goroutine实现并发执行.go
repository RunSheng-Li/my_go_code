package main

import (
	"fmt"
	"time"
)

func haha() {
	time.Sleep(time.Second * 2)
	fmt.Println("劳资睡眠结束了哈哈哈")
}

func main() {

	go haha()
	fmt.Println("我会马上执行")
	time.Sleep(time.Second * 3)

}
