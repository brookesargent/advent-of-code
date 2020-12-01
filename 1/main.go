package main

import (
	"fmt"
	"github.com/brookesargent/advent-of-code2020/helper"
	"log"
	"strconv"
)
func main() {
	lines, err := helper.ReadInputTxtToIntSlice("1/input.txt")
	if err != nil {
		log.Println(err)
	}

	subset1 := subsetSumPartOne(lines, len(lines), 2020)
	subset2 := subsetSumPartTwo(lines, len(lines), 2020)
	fmt.Println("The answer is: " + strconv.Itoa(subset1[0] * subset1[1]))
	fmt.Println("The answer is: " + strconv.Itoa(subset2[0] * subset2[1] * subset2[2]))

}

func subsetSumPartOne(numbers []int, setLength int, sum int) []int {
	var subset []int
	for i := 0; i < setLength; i++ {
		for j:= 1; j < setLength; j++ {
			if numbers[i] +  numbers[j] == sum {
				subset = append(subset, numbers[i], numbers[j])
			}
		}
	}

	return subset
}

func subsetSumPartTwo(numbers []int, setLength int, sum int) []int {
	var subset []int
	for i := 0; i < setLength; i++ {
		for j:= 1; j < setLength; j++ {
			for k := 2; k < setLength; k++ {
				if numbers[i] +  numbers[j] + numbers[k] == sum {
					subset = append(subset, numbers[i], numbers[j], numbers[k])
				}
			}
		}
	}

	return subset
}