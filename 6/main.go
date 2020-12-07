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

type kv struct {
	Key   string
	Value int
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("6/input.txt")
	if err != nil {
		log.Println(err)
	}
	lines = append(lines, "")

	totalYesCountA := totalYesesPart1(lines)
	totalYesCountB := totalYesesPart2(lines)
	fmt.Printf("Part 1 yes count: %s\n", strconv.Itoa(totalYesCountA))
	fmt.Printf("Part 2 yes count: %s\n", strconv.Itoa(totalYesCountB))
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

func totalYesesPart2(lines []string) int {
	participantCount := 0
	totalYesCount := 0
	yesTally := make(map[string]int)
	for _, line := range lines {
		if line == "" {
			// which characters in the map have a value == participant count?
			sortedTally := sortSliceByMapValues(yesTally)
			for i := 0; i < len(sortedTally); i++ {
				if sortedTally[i].Value == participantCount {
					totalYesCount++
				}
			}
			//zero out counters
			yesTally = make(map[string]int)
			participantCount = 0
			continue
		}
		yeses := strings.Split(line, "")
		for _, yes := range yeses {
			// add character to map/increment count
			yesTally[yes] += 1
		}
		participantCount++
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

func sortSliceByMapValues(slice map[string]int) []kv {
	var ss []kv
	for k, v := range slice {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}
