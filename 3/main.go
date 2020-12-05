package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()
	position := 0
	trees := 0
	lines, err := helper.ReadInputTxtToStringSlice("3/input.txt")
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(lines); i++ {
		if i > 0 {
			row := strings.Split(lines[i], "")
			if row[position] == "#" {
				trees++
			}
		}

		switch (len(lines[i]) - 1) - position {
		case 0:
			position = 2
		case 1:
			position = 1
		case 2:
			position = 0
		default:
			position += 3
		}

		continue
	}

	fmt.Printf("There are %s trees\n", strconv.Itoa(trees))
	fmt.Println("Program duration: " + time.Since(start).String())
}
