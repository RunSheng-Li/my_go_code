package main

import (
	"fmt"
	"time"
)

//全局定义channel,用来完成数据同步
var channel = make(chan int)

//定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}

}

func person1() {
	printer("hello")
	channel <- 8
}

func person2() {
	<-channel

	printer("world")

}

func main() {

	go person1()
	go person2()
	for {
		;
	}
}
