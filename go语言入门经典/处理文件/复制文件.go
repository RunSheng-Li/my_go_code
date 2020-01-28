package main

import (
	"io"
	"log"
	"os"
)

func main() {

	from, err := os.Open("./1.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer from.Close()

	to, err := os.OpenFile("./01_copy.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

}

//使用os包中的函数Open来读取磁盘文件
//使用defer语句在程序完成其他所有操作后关闭文件
//使用OpenFile打开文件。第一个参数是要打开（如果不存在，就创建）的文件的名称，第二个参数是用于文件的标志，在这里指定的是读写文件，并在文件不存在时创建它；最后一个参数设置文件的权限
//再次使用defer语句在执行完其他操作后关闭文件
//使用io包中的函数Copy复制源文件的内容，并将其写入目标文件
