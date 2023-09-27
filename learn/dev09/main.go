package main

import "fmt"

/*
Сколько существует способов задать переменную типа slice или map?

new
make
:=
var
*/
func main() {
	w := new([]string)
	qw := make([]string, 0)
	qwe := []string{}
	var qwer []string
	fmt.Println(w, qw, qwe, qwer)
}
