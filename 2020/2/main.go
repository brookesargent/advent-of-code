package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

type PasswordPolicy struct {
	Letter        string
	MinimumCount  int
	MaxiumumCount int
	Password      string
}

func main() {
	start := time.Now()
	lines, err := helper.ReadInputTxtToStringSlice("2/input.txt")
	if err != nil {
		log.Println(err)
	}

	var countA = 0
	var countB = 0
	passwordPolicies := parsePasswordPolicies(lines)
	for _, policy := range passwordPolicies {
		if isPasswordValidCharacterCount(policy) {
			countA++
		}

		if isPasswordValidPosition(policy) {
			countB++
		}
	}
	fmt.Println("Number of valid passwords for part 1: " + strconv.Itoa(countA))
	fmt.Println("Number of valid passwords for part 2: " + strconv.Itoa(countB))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func parsePasswordPolicies(lines []string) []PasswordPolicy {
	var policies []PasswordPolicy
	for _, line := range lines {
		var policy PasswordPolicy
		splitLine := strings.Split(line, " ")
		minMax := strings.Split(splitLine[0], "-")
		policy.MinimumCount, _ = strconv.Atoi(minMax[0])
		policy.MaxiumumCount, _ = strconv.Atoi(minMax[1])
		policy.Letter = strings.Trim(splitLine[1], ":")
		policy.Password = splitLine[2]
		policies = append(policies, policy)
	}
	return policies
}

func isPasswordValidCharacterCount(policy PasswordPolicy) bool {
	if strings.Count(policy.Password, policy.Letter) >= policy.MinimumCount && strings.Count(policy.Password, policy.Letter) <= policy.MaxiumumCount {
		return true
	}

	return false
}

func isPasswordValidPosition(policy PasswordPolicy) bool {
	splitPassword := strings.Split(policy.Password, "")
	return (splitPassword[policy.MinimumCount-1] == policy.Letter) != (splitPassword[policy.MaxiumumCount-1] == policy.Letter)
}
