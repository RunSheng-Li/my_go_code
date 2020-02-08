package main

import "fmt"

//在访问的key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在

func main() {
	m1 := map[int]int{}
	fmt.Println(m1[1])
	m1[2] = 0
	fmt.Println(m1[2])
	//m1[3] = 0
	if v, ok := m1[3]; ok {
		fmt.Println("存在这个key，值为", v)
	} else {
		fmt.Println("不存在这个key")
	}
}

//可以用map实现set相同的特性