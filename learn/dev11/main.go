package main

import (
	"fmt"
	"sync"
)

/*
Что выведет данная программа и почему?

deadlock из за передачи копии wg в горутины.

Если передадим указатель wg, все ок
*/
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
