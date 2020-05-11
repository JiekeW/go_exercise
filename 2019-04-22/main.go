package main

import (
	"fmt"
)

func printHelloWorld() {
	fmt.Println("Hello, World!")
}

func pointer() {
	var a = 3.1415926
	arr := [...]int{1:0}
	fmt.Println("chongchong value a = ", a)
	fmt.Println("chongchong Address a = ", &a)
	fmt.Println("chongchong value p = ", *&a)
	fmt.Println("chongchong value arr = ", arr)
}

type rect struct {
	width int
	height int
}

func (r *rect) area() int {
	fmt.Println(r, *r, &r)
	r.width = 20
	return r.width * r.height
}

func main() {
	printHelloWorld()

	pointer()

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println(r.width)
}