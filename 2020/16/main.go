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

var validTicketRules map[int][]string

type kv struct {
	Key   int
	Value []string
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("16/input.txt")
	if err != nil {
		log.Println(err)
	}

	ticketRules, myTicket, nearbyTickets := parseTicketInfo(lines)
	ticketCategories := make([]string, 0)
	for _, tr := range ticketRules {
		splitRule := strings.Split(tr, ":")
		category := splitRule[0]
		ticketCategories = append(ticketCategories, category)
		values := strings.Split(splitRule[1], "or")
		for _, v := range values {
			determineValidValuesForCategory(category, v)
		}
	}

	errorRate := 0
	validNearbys := make([]string, 0)
	for _, nt := range nearbyTickets {
		valid := true
		ticketValues := strings.Split(nt, ",")
		var num int
		for _, tv := range ticketValues {
			num, err = strconv.Atoi(tv)
			if err != nil {
				fmt.Println(err)
			}
			if _, ok := validTicketRules[num]; !ok {
				errorRate += num
				valid = false
			}
		}
		if valid {
			validNearbys = append(validNearbys, nt)
		}
	}

	validPositions := validSeatPositions(ticketCategories)
	for _, vn := range validNearbys {
		splitNearbys := strings.Split(vn, ",")
		for c, v := range splitNearbys {
			var ticketNum int
			ticketNum, err = strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			// remove any category values from validSeatPositions that aren't present in validTicketRules
			validPositions[c] = removeValues(validPositions[c], validTicketRules[ticketNum])
		}
	}

	sortedKVs := sortMapByValue(validPositions)
	product := 1
	for i, v := range sortedKVs {
		if len(v.Value) == 1 {
			if strings.Contains(v.Value[0], "departure") {
				var n int
				n, err = strconv.Atoi(myTicket[v.Key])
				if err != nil {
					fmt.Println(err)
				}
				product *= n
			}
			for j := i + 1; j < len(sortedKVs); j++ {
				idx := contains(sortedKVs[j].Value, v.Value[0])
				if idx > -1 {
					sortedKVs[j].Value = remove(sortedKVs[j].Value, idx)
				}
			}
		}
	}

	fmt.Println(product)
	fmt.Println(errorRate)
	fmt.Println("Program duration: " + time.Since(start).String())
}

func parseTicketInfo(lines []string) ([]string, []string, []string) {
	var myTicket string
	rules := make([]string, 0)
	nearbyTickets := make([]string, 0)
	validTicketRules = make(map[int][]string)
	lineBreakIndex := indexOf("", lines)
	rules = append(rules, lines[0:lineBreakIndex]...)
	myTicket = lines[lineBreakIndex+2]
	mySplitTicket := strings.Split(myTicket, ",")
	nearbyTickets = append(nearbyTickets, lines[lineBreakIndex+5:]...)
	return rules, mySplitTicket, nearbyTickets
}

func determineValidValuesForCategory(category string, valueRange string) {
	valueRange = strings.Trim(valueRange, " ")
	split := strings.Split(valueRange, "-")
	lower, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println(err)
	}
	upper, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println(err)
	}

	for i := lower; i <= upper; i++ {
		validTicketRules[i] = append(validTicketRules[i], category)
	}
}

func validSeatPositions(ticketCategories []string) map[int][]string {
	positions := make(map[int][]string)
	for i := 0; i < len(ticketCategories); i++ {
		tmp := make([]string, len(ticketCategories))
		copy(tmp, ticketCategories)
		positions[i] = tmp
	}
	return positions
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func removeValues(slice1 []string, slice2 []string) []string {
	for i, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			slice1 = remove(slice1, i)
		}
	}

	return slice1
}

func remove(a []string, i int) []string {
	a = append(a[:i], a[i+1:]...)
	return a
}

func sortMapByValue(m map[int][]string) []kv {
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return len(ss[i].Value) < len(ss[j].Value)
	})

	return ss
}

func contains(slice []string, s string) int {
	for idx, v := range slice {
		if v == s {
			return idx
		}
	}
	return -1
}
