package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {

	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 当前时间
	now = time.Now()
	fmt.Println(now)
	// 下次调度时间
	nextTime = expr.Next(now)
	fmt.Println(nextTime)

	fmt.Println(nextTime.Sub(now))

	// 等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了:", nextTime)
	})

	time.Sleep(5 * time.Second)

}
