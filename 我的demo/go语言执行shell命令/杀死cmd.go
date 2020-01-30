package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {

	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)

	// 创建一个结果队列
	resultChan = make(chan *result, 1000)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {

		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "echo 我开始输出了;sleep 3;echo 跨越长城，走向世界")

		// 执行任务，捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出的结果，传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}

	}()

	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	// 在main协程里，等待子协程的退出，并打印任务执行结果
	res = <-resultChan

	fmt.Println(res.err, string(res.output))

}
