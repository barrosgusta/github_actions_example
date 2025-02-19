package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // Delay for readability
	}
}
