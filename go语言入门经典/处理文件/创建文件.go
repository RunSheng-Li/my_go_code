package main

import (
	"io/ioutil"
	"log"
)

// 在文件系统中创建了一个文件，并将其权限设置为0644
func main() {
	b := make([]byte, 0)
	err := ioutil.WriteFile("1.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
