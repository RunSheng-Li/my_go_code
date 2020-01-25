//常量指的是在整个程序声明周期内都不变的值。常量初始化后，可以引用它，但不能修改它
package main

import "fmt"

const greeting string = "hello,world"

func main() {
	fmt.Println(greeting)
}
//试图修改常量将导致报错