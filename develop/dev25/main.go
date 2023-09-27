package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/
const duration = time.Second * 5

func main() {

	fmt.Println("start")
	sleep(duration)
	fmt.Println("finish")

}
func sleep(d time.Duration) {
	time := time.After(d)
	ctx, cancel := context.WithDeadline(context.Background(), <-time)

	for {
		select {
		case <-ctx.Done():
			cancel()
			return
		}
	}
}
