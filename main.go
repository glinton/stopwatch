package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	curr := time.Now()

	q := make(chan os.Signal, 1)
	signal.Notify(q)

	for {
		select {
		case <-q:
			fmt.Print("\r\033[K")
			fmt.Println(time.Since(curr))
			return
		default:
			fmt.Printf("\r%s", time.Since(curr))
			time.Sleep(100 * time.Millisecond)
		}
	}
}
