package main

import (
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
			os.Stdout.WriteString("\r\033[K" + time.Since(curr).String() + "\n")
			return
		default:
			os.Stdout.WriteString("\r\033[K" + time.Since(curr).String())
			time.Sleep(100 * time.Millisecond)
		}
	}
}
