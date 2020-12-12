package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("11/input.txt")
	if err != nil {
		log.Println(err)
	}

	startLayout := make([][]string, len(lines))
	currentLayout := make([][]string, len(lines))
	for i, line := range lines {
		seats := strings.Split(line, "")
		seats2 := strings.Split(line, "")
		startLayout[i] = seats
		currentLayout[i] = seats2
	}

	for {
		prevTotalOccupied := 0
		totalOccupied := 0
		for outer := 0; outer < len(startLayout); outer++ {
			for inner := 0; inner < len(startLayout[outer]); inner++ {
				adjacentOccupied := 0

				if outer > 0 {
					if checkSeat(startLayout, "up", outer, inner) {
						adjacentOccupied++
					}
				}

				if outer < len(startLayout)-1 {
					if checkSeat(startLayout, "down", outer, inner) {
						adjacentOccupied++
					}
				}

				if inner > 0 {
					if checkSeat(startLayout, "left", outer, inner) {
						adjacentOccupied++
					}
				}

				if inner < len(startLayout[outer])-1 {
					if checkSeat(startLayout, "right", outer, inner) {
						adjacentOccupied++
					}
				}

				// upper left diag
				if outer > 0 && inner > 0 {
					if checkSeat(startLayout, "ul_diag", outer, inner) {
						adjacentOccupied++
					}
				}

				// upper right diag
				if outer > 0 && inner < len(startLayout[outer])-1 {
					if checkSeat(startLayout, "ur_diag", outer, inner) {
						adjacentOccupied++
					}
				}

				// lower left diag
				if outer < len(startLayout)-1 && inner > 0 {
					if checkSeat(startLayout, "ll_diag", outer, inner) {
						adjacentOccupied++
					}
				}

				// lower right diag
				if outer < len(startLayout)-1 && inner < len(startLayout[outer])-1 {
					if checkSeat(startLayout, "lr_diag", outer, inner) {
						adjacentOccupied++
					}
				}

				if startLayout[outer][inner] == "L" && adjacentOccupied == 0 {
					currentLayout[outer][inner] = "#"
					totalOccupied++
				} else if startLayout[outer][inner] == "#" && adjacentOccupied == 4 {
					if adjacentOccupied >= 4 {
						currentLayout[outer][inner] = "L"
					}
					totalOccupied++
				}
			}
		}
		startLayout = currentLayout
		if totalOccupied == prevTotalOccupied {
			fmt.Println(totalOccupied)
			break
		}
	}
	fmt.Println("Program duration: " + time.Since(start).String())
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
