package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("5/input.txt")
	if err != nil {
		log.Println(err)
	}

	var seatIDs []int
	for _, line := range lines {
		columnMin, rowMin := 0, 0
		columnMax, rowMax := 7, 127
		directions := strings.Split(line, "")
		var row int
		var column int
		for i, direction := range directions {
			if i < 7 {
				rowMin, rowMax = calculateRows(rowMin, rowMax, direction)
				if i == 6 {
					if direction == "F" {
						row = rowMin
					} else if direction == "B" {
						row = rowMax
					}
				}
			} else {
				columnMin, columnMax = calculateColumns(columnMin, columnMax, direction)
				if i == 9 {
					if direction == "L" {
						column = columnMin
					} else if direction == "R" {
						column = columnMax
					}
				}
			}
		}
		seatIDs = append(seatIDs, row*8+column)
	}

	sort.Ints(seatIDs)
	mySeat := findMySeat(seatIDs)
	fmt.Printf("Part one: highest seat ID: %s\n", strconv.Itoa(seatIDs[len(seatIDs)-1]))
	fmt.Printf("Part two: my seat ID: %s\n", strconv.Itoa(mySeat))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func calculateRows(min int, max int, direction string) (int, int) {
	var newMin, newMax int
	totalRows := (max - min) + 1
	rowsToTake := totalRows / 2

	if direction == "F" {
		newMin = min
		newMax = (min + rowsToTake) - 1
	} else if direction == "B" {
		newMin = (max - rowsToTake) + 1
		newMax = max
	}

	return newMin, newMax
}

func calculateColumns(min int, max int, direction string) (int, int) {
	var newMin, newMax int
	totalColumns := (max - min) + 1
	columnsToTake := totalColumns / 2

	if direction == "L" {
		newMin = min
		newMax = (min + columnsToTake) - 1
	} else if direction == "R" {
		newMin = (max - columnsToTake) + 1
		newMax = max
	}
	return newMin, newMax
}

func findMySeat(seatIDs []int) int {
	var seatID int
	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i-1] == seatIDs[i]-1 {
			continue
		}
		seatID = seatIDs[i] - 1
		break
	}
	return seatID
}
