package main

import (
	"fmt"
	"time"
)

func haha(c chan string) {

	time.Sleep(time.Second * 2)
	c <- "哈哈哈哈哈哈哈哈"

}

func main() {
	c := make(chan string)
	go haha(c)

	msg := <-c
	fmt.Println(msg)

}

/*
c <- "哈哈哈哈哈哈哈哈"
<-,表示将右边的字符串发给左边的通道。如果通道被指定为收发字符串，则只能向它发送字符串消息，如果向它发送其他类型的消息将导致错误
从通道那里接受消息的语法
msg := <-c


demo解读：
创建一个存储字符串数据的通道，并将其赋值给变量c
函数haha将通道当做参数
haha函数的单个参数指定了一个通道和一个字符串的数据类型
声明变量msg，用于接受来自通道c的消息。这将阻塞进程直到收到消息为止，从而避免进程过早退出
函数haha执行完毕后向通道c发送一条消息
接受并打印这条消息
由于没有其他语句，因此程序就此退出




*/
