package main

import "fmt"

/*
при append переполняется исходный массив, создается новый массив в 2 раза больше первого, копируются старые элементы и добавляется новый.
Получаем слайс len()=3,cap()=4
меняем элементы уже в новом массиве
*/
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
