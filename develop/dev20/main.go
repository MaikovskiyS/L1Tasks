package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func main() {

	str := "snow dog sun"
	newstr := strings.Split(str, " ")
	left := 0
	right := len(newstr) - 1
	for i := 0; i <= (len(newstr)-1)/2; i++ {
		newstr[left], newstr[right] = newstr[right], newstr[left]
		left++
		right--
	}
	fmt.Println(newstr)
}
