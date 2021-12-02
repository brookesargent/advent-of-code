package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
)

var shipDirection string
var x, y, wx, wy, sx, sy int

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
	wx, wy = 10, 1
	instructions := formatInstructions(lines)
	manhattanDistance := partOne(instructions)
	manhattanDistance2 := partTwo(instructions)

	fmt.Println("part 1: " + strconv.Itoa(manhattanDistance))
	fmt.Println("part 2: " + strconv.Itoa(manhattanDistance2))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func partOne(instructions []Instruction) int {
	for _, instruction := range instructions {
		move(instruction.Direction, instruction.Moves)
	}

	manhattanDistance := helper.Abs(x) + helper.Abs(y)
	return manhattanDistance
}

func partTwo(instructions []Instruction) int {
	for _, instruction := range instructions {
		switch instruction.Direction {
		case "N":
			wy += instruction.Moves
		case "S":
			wy -= instruction.Moves
		case "E":
			wx += instruction.Moves
		case "W":
			wx -= instruction.Moves
		case "L":
			rotateWaypoint(instruction.Direction, instruction.Moves)
		case "R":
			rotateWaypoint(instruction.Direction, instruction.Moves)
		case "F":
			moveShipToWaypoint(instruction.Moves)
		default:
			fmt.Println("couldn't move")
		}
	}
	manhattanDistance := helper.Abs(sx) + helper.Abs(sy)
	return manhattanDistance
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

func moveShipToWaypoint(moves int) {
	sx += moves * wx
	sy += moves * wy
}

func rotateWaypoint(direction string, degrees int) {
	turns := degrees / 90
	for i := 0; i < turns; i++ {
		if direction == "L" {
			wx, wy = wy*-1, wx
		} else if direction == "R" {
			wx, wy = wy, wx*-1
		}
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
