package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	input := "14,3,1,0,9,5"
	inputSlice := strings.Split(input, ",")

	a := findNthNumberSpoken(2020, inputSlice)
	b := findNthNumberSpoken(30000000, inputSlice)

	fmt.Println("Part 1 answer: " + strconv.Itoa(a))
	fmt.Println("Part 2 answer: " + strconv.Itoa(b))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func findNthNumberSpoken(n int, input []string) int {
	spoken := make(map[int][]int)

	for i, v := range input {
		n, _ := strconv.Atoi(v)
		spoken[n] = append(spoken[n], i+1)
	}
	prevSpoken := 0
	for j := len(input) + 1; j <= n; j++ {
		var speak int

		if len(spoken[prevSpoken]) < 2 {
			speak = 0
		} else {
			speak = spoken[prevSpoken][1] - spoken[prevSpoken][0]
		}
		if len(spoken[speak]) == 2 {
			// pop the array
			spoken[speak] = pop(spoken[speak])
		}
		spoken[speak] = append(spoken[speak], j)
		prevSpoken = speak
	}
	return prevSpoken
}

func pop(a []int) []int {
	_, a = a[0], a[1:]
	return a
}
