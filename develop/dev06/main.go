package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})
	wg.Add(1)
	go func(chan struct{}, context.Context) {
		defer wg.Done()
		counter := 0
		for {
			counter++
			var d struct{}
			fmt.Println("g1 stopped after writing in chan")
			time.Sleep(time.Second)
			ch <- d
			fmt.Println("g1 do next work")
			if counter == 4 {
				close(ch)
				cancel()
				break
			}
		}
		fmt.Println("g1 work end")
	}(ch, ctx)

	wg.Add(1)
	go func(context.Context) {
		defer wg.Done()
		fmt.Println("g2 stopped by select")
		select {
		case <-ctx.Done():
			fmt.Println("g2 select end")
			break
		}
		fmt.Println("g2 do next work")
		fmt.Println("g2 end ")

	}(ctx)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("g3 stopped for reading from ch")
		for v := range ch {
			_ = v
			time.Sleep(3 * time.Second)
		}
		fmt.Println("g3 do next work")
		fmt.Println("g3 work end")
	}()
	fmt.Println("main stopped by wg")
	wg.Wait()
	fmt.Println(" main do next work")
	fmt.Println("main end ")

}
