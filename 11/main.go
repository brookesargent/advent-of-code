package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

var startLayout [][]string

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("11/input.txt")
	if err != nil {
		log.Println(err)
	}

	startLayout = make([][]string, len(lines))
	currentLayout := make([][]string, len(lines))
	for i, line := range lines {
		seats := strings.Split(line, "")
		seats2 := strings.Split(line, "")
		startLayout[i] = seats
		currentLayout[i] = seats2
	}

	var currentTotalOccupied int
	prevTotalOccupied := 0
	for {
		currentTotalOccupied = 0
		for outer := 0; outer < len(startLayout); outer++ {
			for inner := 0; inner < len(startLayout[outer]); inner++ {
				adjacentOccupied := calcAdjacentOccupied(outer, inner)
				if startLayout[outer][inner] == "L" && adjacentOccupied == 0 {
					currentLayout[outer][inner] = "#"
					currentTotalOccupied++
				} else if startLayout[outer][inner] == "#" {
					if adjacentOccupied >= 4 {
						currentLayout[outer][inner] = "L"
					} else {
						currentTotalOccupied++
					}
				}
			}
		}
		if currentTotalOccupied == prevTotalOccupied {
			break
		}
		startLayout = copyLayout(currentLayout)
		prevTotalOccupied = currentTotalOccupied
	}
	fmt.Println("Total occupied seats: " + strconv.Itoa(currentTotalOccupied))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func calcAdjacentOccupied(outerIdx int, innerIdx int) int {
	adjacentOccupied := 0

	if outerIdx > 0 {
		if checkSeat(startLayout, "up", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	if outerIdx < len(startLayout)-1 {
		if checkSeat(startLayout, "down", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	if innerIdx > 0 {
		if checkSeat(startLayout, "left", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	if innerIdx < len(startLayout[outerIdx])-1 {
		if checkSeat(startLayout, "right", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	// upper left diag
	if outerIdx > 0 && innerIdx > 0 {
		if checkSeat(startLayout, "ul_diag", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	// upper right diag
	if outerIdx > 0 && innerIdx < len(startLayout[outerIdx])-1 {
		if checkSeat(startLayout, "ur_diag", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	// lower left diag
	if outerIdx < len(startLayout)-1 && innerIdx > 0 {
		if checkSeat(startLayout, "ll_diag", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}

	// lower right diag
	if outerIdx < len(startLayout)-1 && innerIdx < len(startLayout[outerIdx])-1 {
		if checkSeat(startLayout, "lr_diag", outerIdx, innerIdx) {
			adjacentOccupied++
		}
	}
	return adjacentOccupied
}

func checkSeat(layout [][]string, direction string, outerIdx int, innerIdx int) bool {
	var isOccupied bool
	switch direction {
	case "up":
		if layout[outerIdx-1][innerIdx] == "#" {
			isOccupied = true
			break
		}
	case "down":
		if layout[outerIdx+1][innerIdx] == "#" {
			isOccupied = true
			break
		}
	case "right":
		if layout[outerIdx][innerIdx+1] == "#" {
			isOccupied = true
			break
		}
	case "left":
		if layout[outerIdx][innerIdx-1] == "#" {
			isOccupied = true
			break
		}
	case "ur_diag":
		if layout[outerIdx-1][innerIdx+1] == "#" {
			isOccupied = true
			break
		}
	case "ul_diag":
		if layout[outerIdx-1][innerIdx-1] == "#" {
			isOccupied = true
			break
		}
	case "lr_diag":
		if layout[outerIdx+1][innerIdx+1] == "#" {
			isOccupied = true
			break
		}
	case "ll_diag":
		if layout[outerIdx+1][innerIdx-1] == "#" {
			isOccupied = true
			break
		}
	}
	return isOccupied
}

func copyLayout(layout [][]string) [][]string {
	newLayout := make([][]string, len(layout))
	for i, value := range layout {
		newSlice := make([]string, len(value))
		copy(newSlice, value)
		newLayout[i] = newSlice
	}
	return newLayout
}
