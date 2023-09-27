package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
const timeout = time.Second * 2

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	counter := Counter{
		value: 0,
	}
	go func(context.Context, Counter) {

		for {
			select {
			case <-ctx.Done():
				return
			default:
				counter.value++
			}
		}
	}(ctx, counter)
	<-ctx.Done()
	cancel()
	fmt.Println(counter)
}

type Counter struct {
	value int
}
