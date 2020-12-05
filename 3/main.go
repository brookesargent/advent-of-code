package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

type Slope struct {
	Right int
	Down  int
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("3/input.txt")
	if err != nil {
		log.Println(err)
	}

	// part one
	slopeA := Slope{
		Right: 3,
		Down:  1,
	}
	treesA := calculateSlope(lines, slopeA)

	// part two
	var treesB []int
	slopeB := []Slope{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	for _, slope := range slopeB {
		treesB = append(treesB, calculateSlope(lines, slope))
	}

	product := treesB[0] * treesB[1] * treesB[2] * treesB[3] * treesB[4]

	fmt.Printf("Part 1: There are %s trees\n", strconv.Itoa(treesA))
	fmt.Printf("Part 2: There are %s trees\n", strconv.Itoa(product))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func calculateSlope(lines []string, slope Slope) int {
	xPosition := 0
	yPosition := 0
	trees := 0
	for i := 0; i < len(lines); i++ {
		if i > 0 {
			// check for a tree if we have traversed down enough, otherwise keep traversing
			if yPosition == slope.Down {
				row := strings.Split(lines[i], "")
				if row[xPosition] == "#" {
					trees++
				}
				yPosition = 0
			} else {
				yPosition++
				continue
			}
		}

		//move right
		rightMoves := (len(lines[i]) - 1) - xPosition
		if rightMoves < slope.Right {
			xPosition = (slope.Right - 1) - rightMoves
		} else {
			xPosition += slope.Right
		}

		//move down
		yPosition++
		continue
	}

	return trees
}
