package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
func main() {
	arr := []int{4, 2, 3, 6, 8, 1, 9, 5}
	result := quickSort(arr)
	fmt.Println(result)
}
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}

	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}
