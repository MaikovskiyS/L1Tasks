package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
func main() {
	xch := make(chan int)
	qch := make(chan float64)
	xarr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(chan int, []int) {
		defer wg.Done()
		for _, number := range xarr {
			xch <- number
		}
		close(xch)
	}(xch, xarr)
	wg.Add(1)
	go func(<-chan int, chan float64) {
		defer wg.Done()
		for number := range xch {
			res := math.Pow(float64(number), 2)
			qch <- res
		}
		close(qch)
	}(xch, qch)
	for v := range qch {
		fmt.Println(v)
	}
	wg.Wait()
}
