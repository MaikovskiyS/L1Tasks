package main

import "fmt"

/*
создаем новую переменную в области видимости if.
*/
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
