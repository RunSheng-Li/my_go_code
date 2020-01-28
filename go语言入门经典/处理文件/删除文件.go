package main

import (
	"log"
	"os"
)

func main() {

	err := os.Remove("1.txt")
	if err != nil {
		log.Fatal(err)
	}

}
//使用这个函数时，不会发出警告，也无法将删除的文件恢复，因为务必要谨慎
