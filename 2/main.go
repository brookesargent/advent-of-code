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

	var count = 0
	passwordPolicies := parsePasswordPolicies(lines)
	for _, policy := range passwordPolicies {
		if isPasswordValid(policy) {
			count++
		}
	}
	fmt.Println("Valid passwords: " + strconv.Itoa(count))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func parsePasswordPolicies(lines []string) []PasswordPolicy {
	var policies []PasswordPolicy
	for _, v := range lines {
		var policy PasswordPolicy
		splitLine := strings.Split(v, ":")
		policy.Password = splitLine[1]
		policyRule := splitLine[0]
		splitRule := strings.Split(policyRule, "-")
		policy.MinimumCount, _ = strconv.Atoi(splitRule[0])
		lastSplit := strings.Split(splitRule[1], " ")
		policy.MaxiumumCount, _ = strconv.Atoi(lastSplit[0])
		policy.Letter = lastSplit[1]
		policies = append(policies, policy)
	}
	return policies
}

func isPasswordValid(policy PasswordPolicy) bool {
	letterCount := 0
	splitPassword := strings.Split(policy.Password, "")
	for _, v := range splitPassword {
		if v == policy.Letter {
			letterCount++
		}
	}

	if letterCount >= policy.MinimumCount && letterCount <= policy.MaxiumumCount {
		return true
	}
	return false
}
