package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
)

type BagRule struct {
	BagCount int
	BagType  string
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("7/input.txt")
	if err != nil {
		log.Println(err)
	}

	bagRules := formatBagRules(lines)
	goldBagHolderCount := countGoldBagHolders(bagRules)
	goldBagContainsCount := countBagsGoldContains(bagRules)
	fmt.Printf("%s bags could contain a gold bag\n", strconv.Itoa(goldBagHolderCount))
	fmt.Printf("Gold bags contain %s other bags\n", strconv.Itoa(goldBagContainsCount))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func countGoldBagHolders(bagRules map[string][]BagRule) int {
	var allHolders []string
	directHolders := getDirectBagHolders(bagRules)
	allHolders = append(allHolders, directHolders...)
	indirectHolders := getIndirectBagHolders(bagRules, directHolders)
	allHolders = append(allHolders, indirectHolders...)
	for len(indirectHolders) != 0 {
		indirectHolders = getIndirectBagHolders(bagRules, indirectHolders)
		allHolders = append(allHolders, indirectHolders...)
	}
	allHolders = helper.RemoveDuplicateValues(allHolders)
	return len(allHolders)
}

func countBagsGoldContains(bagRules map[string][]BagRule) int {
	count := 0
	queue := []BagRule{{
		BagCount: 1,
		BagType:  "shiny gold",
	}}
	for len(queue) > 0 {
		rule := bagRules[queue[0].BagType]
		for _, r := range rule {
			factor := queue[0].BagCount * r.BagCount
			count += factor
			queue = append(queue, BagRule{factor, r.BagType})
		}
		queue = pop(queue)
	}
	return count
}

func getDirectBagHolders(bagRules map[string][]BagRule) []string {
	var directHolders []string

	for bag, rules := range bagRules {
		for _, rule := range rules {
			if rule.BagType == "shiny gold" {
				directHolders = append(directHolders, bag)
			}
		}
	}
	return directHolders
}

func getIndirectBagHolders(bagRules map[string][]BagRule, confirmedHolders []string) []string {
	var indirectHolders []string

	sort.Slice(confirmedHolders, func(i, j int) bool {
		return confirmedHolders[i] <= confirmedHolders[j]
	})

	for bag, rules := range bagRules {
		for _, rule := range rules {
			idx := sort.SearchStrings(confirmedHolders, rule.BagType)
			if idx > len(confirmedHolders)-1 {
				continue
			}

			if confirmedHolders[idx] == rule.BagType {
				// found an indirect holder
				indirectHolders = append(indirectHolders, bag)
			}
		}
	}
	indirectHolders = helper.RemoveDuplicateValues(indirectHolders)
	return indirectHolders
}

func formatBagRules(lines []string) map[string][]BagRule {
	bagRules := make(map[string][]BagRule)
	for _, line := range lines {
		splitLine := strings.Split(line, " bags contain ")
		var currentRules []BagRule
		key := splitLine[0]

		if strings.Contains(splitLine[1], "no") {
			continue
		}

		splitRule := strings.Split(splitLine[1], ", ")
		for _, rule := range splitRule {
			var bagType []string
			rulePieces := strings.Split(rule, " ")
			count, err := strconv.Atoi(rulePieces[0])
			if err != nil {
				log.Println(err)
			}
			for i := 1; i < len(rulePieces)-1; i++ {
				bagType = append(bagType, rulePieces[i])
			}

			structuredRule := BagRule{
				BagCount: count,
				BagType:  strings.Join(bagType, " "),
			}
			currentRules = append(currentRules, structuredRule)
		}
		bagRules[key] = currentRules
	}
	return bagRules
}

func pop(a []BagRule) []BagRule {
	_, a = a[0], a[1:]
	return a
}
