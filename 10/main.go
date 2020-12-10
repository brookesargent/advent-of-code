package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	numbers, err := helper.ReadInputTxtToIntSlice("10/input.txt")
	if err != nil {
		log.Println(err)
	}

	numbers = append(numbers, 0)
	sort.Ints(numbers)

	diffOne, diffThree := calcJoltDiffs(numbers)

	fmt.Println(diffOne * diffThree)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func calcJoltDiffs(numbers []int) (int, int) {
	diffOne, diffThree := 0, 1
	for i := 1; i < len(numbers); i++ {
		voltDiff := numbers[i] - numbers[i-1]
		if voltDiff < 0 {
			continue
		} else if voltDiff == 1 {
			diffOne++
		} else if voltDiff == 3 {
			diffThree++
		}
	}

	return diffOne, diffThree
}
