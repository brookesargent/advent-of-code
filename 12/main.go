package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

var shipDirection string
var x, y int

type Instruction struct {
	Direction string
	Moves     int
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("12/input.txt")
	if err != nil {
		log.Println(err)
	}

	shipDirection = "E"
	x, y = 0, 0
	instructions := formatInstructions(lines)

	for _, instruction := range instructions {
		move(instruction.Direction, instruction.Moves)
	}

	manhattanDistance := helper.Abs(x) + helper.Abs(y)
	fmt.Println(manhattanDistance)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func move(direction string, moves int) {
	switch direction {
	case "N":
		y += moves
	case "S":
		y -= moves
	case "E":
		x += moves
	case "W":
		x -= moves
	case "L":
		turn(direction, moves)
	case "R":
		turn(direction, moves)
	case "F":
		move(shipDirection, moves)
	default:
		fmt.Println("couldn't move")
	}
}

func turn(turnDirection string, degrees int) {
	compass := []byte{'N', 'E', 'S', 'W'}
	turnsToMake := degrees / 90
	currentIndex := bytes.IndexAny(compass, shipDirection)
	if turnDirection == "R" {
		for (currentIndex + turnsToMake) >= len(compass) {
			turnsToMake = turnsToMake - 4
		}
		shipDirection = string(compass[currentIndex+turnsToMake])
	} else if turnDirection == "L" {
		for (currentIndex - turnsToMake) < 0 {
			turnsToMake = turnsToMake - 4
		}
		shipDirection = string(compass[currentIndex-turnsToMake])
	}
}

func formatInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		lineSlice := strings.Split(line, "")
		moves, err := strconv.Atoi(strings.Join(lineSlice[1:], ""))
		if err != nil {
			log.Println(err)
		}
		instructions = append(instructions, Instruction{
			Direction: lineSlice[0],
			Moves:     moves,
		})
	}
	return instructions
}
