package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
type Type interface {
}

func main() {
	var t Type

	t = 0
	fmt.Println(getType(t))

	t = "abc"
	fmt.Println(getType(t))

	t = true
	fmt.Println(getType(t))

	t = make(chan int)
	fmt.Println(getType(t))
}

func getType(t Type) string {
	return fmt.Sprintf("Type is %T", t)
}
