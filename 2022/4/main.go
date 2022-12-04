package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
)

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToStringSlice("input.txt")
	if err != nil {
		log.Println(err)
	}

	contained := 0
	overlap := 0

	for _, v := range lines {
		// part 1
		points := toIntPoints(v)
		if (points[0] <= points[2] && points[1] >= points[3]) || (points[2] <= points[0] && points[3] >= points[1]) {
			contained++
		}

		// part 2
		if (points[0] <= points[3] && points[2] <= points[1]) {
			overlap++
		}
	}

	fmt.Printf("Answer 1 is: %d\n", contained)
	fmt.Printf("Answer 2 is: %d\n", overlap)

	fmt.Println("Program duration: " + time.Since(start).String())
}

func toIntPoints(s string) []int {
	points := make([]int, 0)
	pair := strings.Split(s, ",")
	elf1 := strings.Split(pair[0], "-")
	elf2 := strings.Split(pair[1], "-")

	for _, e1 := range elf1 {
		num, _ := strconv.Atoi(e1)
		points = append(points, num)
	}

	for _, e2 := range elf2 {
		num, _ := strconv.Atoi(e2)
		points = append(points, num)
	}

	return points
}
