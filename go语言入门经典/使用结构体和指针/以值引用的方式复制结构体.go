package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {

	a := Drink{
		Name: "可乐",
		Ice:  true,
	}

	b := a
	b.Ice = false
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)

}

//赋值后，a与b相同，但它是b的副本，而不是指向b的引用。修改b不会影响a，反之亦然
