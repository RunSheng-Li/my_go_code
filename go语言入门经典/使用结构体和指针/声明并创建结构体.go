package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {

	m := Movie{
		Name:   "钢铁侠",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)

}

//要访问结构体的数据字段，可使用点表示法：结构体变量名、圆点和要访问的数据字段的名称