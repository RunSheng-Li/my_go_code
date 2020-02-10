package main

import "fmt"

//三角
type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 4}
	t.changeBase(4)
	fmt.Println(t.base)
}

//之所以打印的是3，是因为方法changeBase接受的是一个值引用。这意味着这个方法操作的是结构体Triangle的副本，而原始结构体不受影响。在方法changeBase中，修改的是原始Triangle的结构体副本的t.base
