package main

import "fmt"

/*
Что выведет данная программа и почему?

в go все передается по значению. в функцию мы передали копию адреса памяти, которая лежит уже на другом адресе.
*/

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	fmt.Println(p)
	update(p)
	fmt.Println(*p)
}
func update(g *int) {
	fmt.Println("in func:", &g)
	b := 2
	g = &b
}
