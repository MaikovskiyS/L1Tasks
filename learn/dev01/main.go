package main

import (
	"fmt"
	"strings"
)

/*
Какой самый эффективный способ конкатенации строк?
*/
func main() {
	strings := []string{"This ", "is ", "string"}

	b := concat2builder(strings)
	fmt.Println(b)
}

// билдер не аллоцирует память под каждую строку
func concat2builder(str []string) string {
	var builder strings.Builder
	for _, v := range str {
		builder.Grow(len(v)) // grow capacity
		builder.WriteString(v)
	}
	return builder.String()
}
