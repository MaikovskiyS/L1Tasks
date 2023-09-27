package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
fhdjswqyt - true
*/
func main() {

	str := "fhdjswqyt"
	fmt.Println(uniqstr(str))

}
func uniqstr(str string) bool {
	mapa := make(map[rune]struct{})
	for _, v := range str {
		_, ok := mapa[v]
		if ok {
			return false
		}
		mapa[v] = struct{}{}
	}
	return true
}
