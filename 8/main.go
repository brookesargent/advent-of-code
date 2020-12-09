package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/brookesargent/advent-of-code2020/helper"
)

type GameMove struct {
	Type      string
	Count     int
	Direction string
	Used      bool
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("8/input.txt")
	if err != nil {
		log.Println(err)
	}
	moves1 := formatGameMoves(lines)
	moves2 := append([]GameMove(nil), moves1...)

	// part 1
	accumulator1, _ := runBootCode(moves1)

	// part 2
	var accumulator2 int
	var programComplete bool
	for i, move := range moves2 {
		currentMoves := append([]GameMove(nil), moves2...)
		if move.Type == "jmp" {
			currentMoves[i].Type = "nop"
			accumulator2, programComplete = runBootCode(currentMoves)
		} else if move.Type == "nop" {
			currentMoves[i].Type = "jmp"
			accumulator2, programComplete = runBootCode(currentMoves)
		}

		if programComplete {
			break
		}
	}

	fmt.Println("Accumulator value: " + strconv.Itoa(accumulator1))
	fmt.Println("Accumulator value: " + strconv.Itoa(accumulator2))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func formatGameMoves(lines []string) []GameMove {
	var moveList []GameMove
	for _, line := range lines {
		moves := strings.Split(line, " ")

		var direction string
		if strings.Contains(moves[1], "+") {
			direction = "+"
		} else {
			direction = "-"
		}

		countInt, err := strconv.Atoi(trimFirstRune(moves[1]))
		if err != nil {
			log.Println(err)
		}
		moveList = append(moveList, GameMove{
			Type:      moves[0],
			Count:     countInt,
			Direction: direction,
			Used:      false,
		})
	}
	return moveList
}

func runBootCode(moves []GameMove) (int, bool) {
	accumulator, idx := 0, 0
	programComplete := false

	for {
		currentMove := moves[idx]
		if currentMove.Used {
			break
		} else {
			moves[idx].Used = true
		}

		if currentMove.Type == "nop" {
			idx++
		} else if currentMove.Type == "acc" {
			if currentMove.Direction == "+" {
				accumulator += currentMove.Count
			} else if currentMove.Direction == "-" {
				accumulator -= currentMove.Count
			}
			idx++
		} else if currentMove.Type == "jmp" {
			if currentMove.Direction == "+" {
				idx += currentMove.Count
			} else if currentMove.Direction == "-" {
				idx -= currentMove.Count
			}
		}

		if idx >= len(moves) {
			programComplete = true
			break
		}
	}

	return accumulator, programComplete
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
