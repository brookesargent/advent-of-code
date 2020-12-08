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
	moveList := formatGameMoves(lines)
	accumulator, idx := 0, 0
	for {
		currentMove := moveList[idx]
		if currentMove.Used {
			break
		} else {
			moveList[idx].Used = true
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
	}
	fmt.Println("Accumulator value: " + strconv.Itoa(accumulator))
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

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
