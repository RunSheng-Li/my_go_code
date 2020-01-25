//空指针：未被初始化的指针。   var p *int     *p--->error

//package main
//
//import "fmt"
//
//func main() {
//
//	var p *int
//	fmt.Println(*p)
//
//}

//野指针：被一片无效的地址空间初始化。

//package main
//
//import "fmt"
//
//func main() {
//	var p *int = 0
//	fmt.Println(*p)
//
//}
