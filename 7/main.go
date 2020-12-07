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
	indirectHolders := countIndirectBagHolders(bagRules, directHolders)
	allHolders = append(allHolders, indirectHolders...)
	for len(indirectHolders) != 0 {
		indirectHolders = countIndirectBagHolders(bagRules, indirectHolders)
		allHolders = append(allHolders, indirectHolders...)
	}
	allHolders = removeDuplicateValues(allHolders)
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
		//push types in rule to queue
		for _, r := range rule {
			factor := queue[0].BagCount * r.BagCount
			count += factor
			queue = append(queue, BagRule{factor, r.BagType})
		}
		//front pop from list
		queue = shift(queue)
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

func countIndirectBagHolders(bagRules map[string][]BagRule, confirmedHolders []string) []string {
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
	indirectHolders = removeDuplicateValues(indirectHolders)
	return indirectHolders
}

func formatBagRules(lines []string) map[string][]BagRule {
	bagRules := make(map[string][]BagRule)
	for _, line := range lines {
		line = strings.Trim(line, ".")
		splitLine := strings.Split(line, "contain")
		var currentRules []BagRule
		key := strings.TrimSuffix(splitLine[0], " bags ")
		if strings.Contains(splitLine[1], "no") {
			structuredRule := BagRule{
				BagCount: 0,
			}
			currentRules = append(currentRules, structuredRule)
			bagRules[key] = currentRules
			continue
		}

		splitRule := strings.Split(strings.Trim(splitLine[1], " "), ",")
		for _, rule := range splitRule {
			var bagType []string
			rulePieces := strings.Split(strings.Trim(rule, " "), " ")
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

func shift(a []BagRule) []BagRule {
	_, a = a[0], a[1:]
	return a
}
