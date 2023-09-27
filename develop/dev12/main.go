package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := SetMany(s)
	fmt.Println(set)
}
func SetMany(s []string) []string {
	resMap := make(map[string]bool)
	result := []string{}
	for _, key := range s {
		resMap[key] = true
	}
	for key := range resMap {
		result = append(result, key)
	}
	return result
}
