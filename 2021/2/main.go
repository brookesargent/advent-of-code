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
	lines, err := helper.ReadInputTxtToStringSlice("2021/2/input.txt")
	if err != nil {
		log.Println(err)
	}

	pos := 0
	depth := 0

	for _, line := range lines {
		directions := strings.Split(line, " ")
		moves, _ := strconv.Atoi(directions[1])
		if directions[0] == "forward" {
			pos += moves
		} else if directions[0] == "down" {
			depth += moves
		} else if directions[0] == "up" {
			depth -= moves
		}
	}

	fmt.Printf("Answer 1 is: %d\n", pos*depth)
	fmt.Println("Program duration: " + time.Since(start).String())
}
