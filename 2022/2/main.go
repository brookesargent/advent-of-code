package main

import (
	"fmt"
	"log"
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

	score1 := 0
	score2 := 0

	for _, v := range lines {
		round := strings.Split(v, " ")

		opponent := unencrypt(round[0])

		// part 1
		me := unencrypt(round[1])
		score1 += score(opponent, me)

		// part 2
		me = chooseShape(opponent, round[1])
		score2 += score(opponent, me)
	}

	fmt.Printf("Answer 1 is: %d\n", score1)
	fmt.Printf("Answer 2 is: %d\n", score2)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func unencrypt(s string) string {
	switch s {
	case "A", "X":
		return "Rock"
	case "B", "Y":
		return "Paper"
	case "C", "Z":
		return "Scissors"
	default:
		return ""
	}
}

func score(opponent, me string) int {
	score := 0

	if opponent == me {
		score += 3
	}

	switch me {
	case "Rock":
		score += 1
		if opponent == "Scissors" {
			score += 6
		}
	case "Paper":
		score += 2
		if opponent == "Rock" {
			score += 6
		}
	case "Scissors":
		score += 3
		if opponent == "Paper" {
			score += 6
		}
	}

	return score
}

func chooseShape(opponent, outcome string) string {
	if outcome == "Y" {	// need to draw
		return opponent
	} else if outcome == "X" {	// need to lose
		switch opponent {
		case "Rock":
			return "Scissors"
		case "Paper":
			return "Rock"
		case "Scissors":
			return "Paper"
		default:
			return ""
		}
	} else if outcome == "Z" {	// need to win
		switch opponent {
		case "Rock":
			return "Paper"
		case "Paper":
			return "Scissors"
		case "Scissors":
			return "Rock"
		default:
			return ""
		}
	}
	return ""
}
