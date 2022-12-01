package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
)

type Elf struct {
	ID       int
	Calories int
}

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToIntSlice("input.txt")
	if err != nil {
		log.Println(err)
	}

	currentElf := 1
	currentCalorieTotal := 0
	elves := make([]Elf, 0)

	for _, v := range lines {
		if v == 0 {
			elves = append(elves, Elf{
				ID:       currentElf,
				Calories: currentCalorieTotal,
			})
			currentElf++
			currentCalorieTotal = 0
		} else {
			currentCalorieTotal += v
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})

	top3Total := elves[0].Calories + elves[1].Calories + elves[2].Calories

	fmt.Printf("Answer 1 is: %d\n", elves[0].Calories)
	fmt.Printf("Answer 2 is: %d\n", top3Total)
	fmt.Println("Program duration: " + time.Since(start).String())
}
