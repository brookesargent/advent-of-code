package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Program duration: " + time.Since(start).String())
}
