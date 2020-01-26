package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie
	//没有给字段赋值，因此它们默认为零值。对于字符串，零值为空字符串"";对于float32,零值为0
	fmt.Printf("%+v\n", m)
	//使用点表示法给结构体的数据字段赋值
	m.Name = "钢铁侠"
	m.Rating = 0.99
	fmt.Printf("%+v\n", m)

}
