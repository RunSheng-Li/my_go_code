//指针的函数传参（传引用）
//传地址（引用）：将形参的地址值作为函数参数传递
//传值（数据值）：将实参的 值 拷贝一份给形参
//传引用：在A栈帧内部，修改B栈帧中的变量值
package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println("swap a:", a, "b:", b)
}

func swap2(x, y *int) {
	*x, *y = *y, *x

}

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Println("swap1:main a:", a, "b", b)

	swap2(&a, &b)
	fmt.Println("swap2:main a:", a, "b", b)

}
