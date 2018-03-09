package main

import (
	"fmt"
	"time"
)

func main() {
	curr := time.Now()
	for {
		fmt.Printf("\r%v", time.Since(curr))
		time.Sleep(100 * time.Millisecond)
	}
}