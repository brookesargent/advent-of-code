package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToIntSlice("1/input.txt")
	if err != nil {
		log.Println(err)
	}

	subset1 := subsetSumPartOne(lines, 2020)
	subset2 := subsetSumPartTwo(lines, 2020)
	fmt.Println("The answer is: " + strconv.Itoa(subset1[0]*subset1[1]))
	fmt.Println("The answer is: " + strconv.Itoa(subset2[0]*subset2[1]*subset2[2]))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func subsetSumPartOne(numbers []int, sum int) []int {
	var subset []int
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers)-1; j++ {
			if numbers[i]+numbers[j] == sum {
				subset = append(subset, numbers[i], numbers[j])
			}
		}
	}
	return subset
}

func subsetSumPartTwo(numbers []int, sum int) []int {
	var subset []int
	for i := 0; i < len(numbers); i++ {
		for j := 1; j < len(numbers)-1; j++ {
			for k := 2; k < len(numbers)-2; k++ {
				if numbers[i]+numbers[j]+numbers[k] == sum {
					subset = append(subset, numbers[i], numbers[j], numbers[k])
				}
			}
		}
	}
	return subset
}
