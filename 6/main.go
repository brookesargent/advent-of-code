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

	lines, err := helper.ReadInputTxtToStringSlice("6/input.txt")
	if err != nil {
		log.Println(err)
	}
	lines = append(lines, "")

	totalYesCountA := totalYesesPart1(lines)
	fmt.Println(totalYesCountA)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func totalYesesPart1(lines []string) int {
	totalYesCount := 0
	var yesTally []string
	for _, line := range lines {
		if line == "" {
			// total group yeses
			yesTally = removeDuplicateValues(yesTally)
			totalYesCount += len(yesTally)
			yesTally = nil
			continue
		}
		yeses := strings.Split(line, "")
		yesTally = append(yesTally, yeses...)
	}
	return totalYesCount
}

func removeDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
