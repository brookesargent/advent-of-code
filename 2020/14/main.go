package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("14/input.txt")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(lines)

	fmt.Println("Program duration: " + time.Since(start).String())
}
