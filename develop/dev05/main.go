package main

import (
	"context"
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

const N = 5 //timeout in seconds

func main() {
	input := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*N)
	defer cancel()
	go func(context.Context, chan int) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx Done")
				close(input)
				return
			default:
				input <- i
			}
			i++
		}
	}(ctx, input)

	for value := range input {
		fmt.Println(value)
		_ = value
	}
	fmt.Println("main end")
}
