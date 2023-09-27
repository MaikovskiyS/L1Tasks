package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	result := intersection(a, b)
	fmt.Println(result)
}
func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}
	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}
