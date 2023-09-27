package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
func main() {

	str := "абырвалг"
	result := reverseStr(str)
	fmt.Println(result)
}
func reverseStr(arr string) string {
	rarr := []rune(arr)
	left := 0
	right := len(rarr) - 1
	for i := 0; left != len(rarr)/2; i++ {
		rarr[left], rarr[right] = rarr[right], rarr[left]
		left++
		right--

	}
	return string(rarr)
}
