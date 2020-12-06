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

	lines, err := helper.ReadInputTxtToStringSlice("4/input.txt")
	if err != nil {
		log.Println(err)
	}
	lines = append(lines, "")
	validPassports := 0
	currentPassport := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			if len(currentPassport) == 8 {
				validPassports++
			} else if len(currentPassport) == 7 {
				_, ok := currentPassport["cid"]
				if !ok {
					validPassports++
				}
			}
			currentPassport = make(map[string]string)
		} else {
			passportInfo := strings.Split(line, " ")
			for _, info := range passportInfo {
				keyValPair := strings.Split(info, ":")
				currentPassport[keyValPair[0]] = keyValPair[1]
			}
		}
	}

	fmt.Printf("There are %s valid passports", strconv.Itoa(validPassports))
	fmt.Println("Program duration: " + time.Since(start).String())
}
