package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(3 * time.Second)
	for i := 0; i < 10; i++ {
		time := <-ticker.C
		fmt.Println(time.String())
	}

}

/*
go是 通过 chan

的阻塞实现的。

调用的地方，读取chan

定时，时间到，向chan写入值，阻塞解除，调用函数。


*/