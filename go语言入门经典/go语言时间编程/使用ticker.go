package main

import (
	"fmt"
	"time"
)

func main() {

	c := time.Tick(5 * time.Second)
	for t := range c {
		fmt.Printf("现在的时间是%v\n", t)
	}

}
