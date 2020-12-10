package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
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
	possibleCombos := calcCombos(numbers)
	fmt.Println("Part 1: " + strconv.Itoa(diffOne*diffThree))
	fmt.Println("Part 2: " + strconv.Itoa(possibleCombos))
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

func calcCombos(numbers []int) int {
	numMap := makeNumMap(numbers)
	paths := make([]int, len(numbers))
	paths[len(numbers)-1] = 1
	for i := len(numbers) - 2; i >= 0; i-- {
		sum := 0
		if pos, ok := numMap[numbers[i]+1]; ok {
			sum += paths[pos]
		}

		if pos, ok := numMap[numbers[i]+2]; ok {
			sum += paths[pos]
		}

		if pos, ok := numMap[numbers[i]+3]; ok {
			sum += paths[pos]
		}
		paths[i] = sum
	}
	return paths[0]
}

func makeNumMap(numbers []int) map[int]int {
	numMap := make(map[int]int)
	for i, v := range numbers {
		numMap[v] = i
	}
	return numMap
}
