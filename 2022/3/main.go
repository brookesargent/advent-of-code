package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
	"golang.org/x/exp/slices"
)

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToStringSlice("input.txt")
	if err != nil {
		log.Println(err)
	}

	// part 1
	common := make([]string, 0)

	for _, v := range lines {
		c1 := v[0:len(v)/2]
		c2 := v[len(v)/2:]
		common = append(common, findCommonItemsIn2(c1, c2)...)
	}

	prioritySum := calculatePriority(common)

	fmt.Printf("Answer 1 is: %d\n", prioritySum)

	// part 2
	common = make([]string, 0)
	for i := 0; i < len(lines); i+=3 {
		commonChar := findCommonItemsIn3(lines[i], lines[i+1], lines[i+2])
		common = append(common, commonChar)
	}

	prioritySum = calculatePriority(common)
	
	fmt.Printf("Answer 2 is: %d\n", prioritySum)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func findCommonItemsIn2(s1, s2 string) []string {
	common := make([]string, 0)
	for i := 0; i < len(s1); i++ {
		char := string(s1[i])
		if strings.Contains(s2, char) && !slices.Contains(common, char) {
			common = append(common, char)
		}
	}
	return common
}

func findCommonItemsIn3(s1, s2, s3 string) string {
	var common string
	for i := 0; i < len(s1); i++ {
		char := string(s1[i])
		if strings.Contains(s2, char) {
			if strings.Contains(s3, char) {
				return char
			}
		}
	}
	return common
}

func calculatePriority(items []string) int {
	sum := 0
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, i := range items {
		sum += strings.Index(alpha, i) + 1
	}
	return sum
}