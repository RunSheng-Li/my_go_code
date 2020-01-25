package main

import "fmt"

func main() {

	var p *int

	//在heap上申请一片内存地址空间

	p = new(int)
	*p = 1000
	fmt.Printf("%v\n", *p)

}
