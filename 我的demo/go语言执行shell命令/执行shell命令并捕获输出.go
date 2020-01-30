package main

import (
	"fmt"
	"os/exec"
)

func main() {

	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成cmd
	cmd = exec.Command("/bin/bash", "-c", "echo 跨越长城，走向世界")

	// 执行命令，捕获子进程的输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))

}
