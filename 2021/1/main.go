package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
)

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToIntSlice("2021/1/input.txt")
	if err != nil {
		log.Println(err)
	}

	incCount := 0
	var prevMeasurement int
	for i, v := range lines {
		if i > 0 && v > prevMeasurement {
			incCount++
		}
		prevMeasurement = v
	}
	fmt.Println(fmt.Sprintf("Answer 1 is: %d", incCount))
	fmt.Println("Program duration: " + time.Since(start).String())
}
