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
	spoken := make(map[int][]int)

	for i, v := range inputSlice {
		n, _ := strconv.Atoi(v)
		spoken[n] = append(spoken[n], i + 1)
	}


	prevSpoken := 0
	for j := len(inputSlice)+1; j <= 2020; j++ {
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
	
	fmt.Println(prevSpoken)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func pop(a []int) []int {
	_, a = a[0], a[1:]
	return a
}