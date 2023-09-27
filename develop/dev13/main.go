package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
func main() {

	a := 1
	b := 2
	fmt.Printf("a:%v b:%v\n", a, b)
	a, b = b, a
	fmt.Printf("a:%v b:%v\n", a, b)
}
