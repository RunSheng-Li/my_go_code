package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))

}

//如果找到，就返回第一个子串的索引号
//如果没找到，就返回-1
