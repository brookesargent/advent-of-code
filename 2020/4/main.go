package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code/helper"
	"github.com/mitchellh/mapstructure"
)

type Passport struct {
	BirthYear      int    `mapstructure:"byr"`
	IssueYear      int    `mapstructure:"iyr"`
	ExpirationYear int    `mapstructure:"eyr"`
	Height         string `mapstructure:"hgt"`
	HairColor      string `mapstructure:"hcl"`
	EyeColor       string `mapstructure:"ecl"`
	PassportID     string `mapstructure:"pid"`
	CountryID      string `mapstructure:"cid"`
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("4/input.txt")
	if err != nil {
		log.Println(err)
	}
	lines = append(lines, "")

	// part one
	validPassportCountA := validatePassportsA(lines)

	// part two
	passports := sanitizePassports(lines)
	fmt.Println(len(passports))
	if err != nil {
		fmt.Println(err)
	}
	validPassportCountB := validatePassportsB(passports)

	fmt.Printf("Part 1: there are %s valid passports", strconv.Itoa(validPassportCountA))
	fmt.Printf("Part 2: there are %s valid passports", strconv.Itoa(validPassportCountB))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func validatePassportsA(lines []string) int {
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
	return validPassports
}

func sanitizePassports(lines []string) []Passport {
	var passports []Passport
	currentPassport := make(map[string]interface{})
	for _, line := range lines {
		if line == "" {
			var p Passport
			err := mapstructure.Decode(currentPassport, &p)
			if err != nil {
				panic(err)
			}
			passports = append(passports, p)
			currentPassport = make(map[string]interface{})
		} else {
			passportInfo := strings.Split(line, " ")
			for _, info := range passportInfo {
				keyValPair := strings.Split(info, ":")
				if keyValPair[0] == "byr" || keyValPair[0] == "iyr" || keyValPair[0] == "eyr" {
					currentPassport[keyValPair[0]], _ = strconv.Atoi(keyValPair[1])
				} else {
					currentPassport[keyValPair[0]] = keyValPair[1]
				}
			}
		}
	}

	return passports
}

func validatePassportsB(passports []Passport) int {
	validPassports := 0
	for _, passport := range passports {
		// validate birth year
		if countDigits(passport.BirthYear) == 4 {
			if !validYear(passport.BirthYear, 1920, 2002) {
				continue
			}
		} else {
			continue
		}

		// validate issue year
		if countDigits(passport.IssueYear) == 4 {
			if !validYear(passport.IssueYear, 2010, 2020) {
				continue
			}
		} else {
			continue
		}

		// validate expiration year
		if countDigits(passport.ExpirationYear) == 4 {
			if !validYear(passport.ExpirationYear, 2020, 2030) {
				continue
			}
		} else {
			continue
		}

		// validate height
		if !validHeight(passport.Height) {
			continue
		}
		// validate hair color
		if !validHair(passport.HairColor) {
			continue
		}
		// validate eye color
		if !validEye(passport.EyeColor) {
			continue
		}
		// validate id
		if !validID(passport.PassportID) {
			continue
		}
		validPassports++
	}
	return validPassports
}

func countDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func validYear(i int, min int, max int) bool {
	if i >= min && i <= max {
		return true
	}
	return false
}

func validHeight(height string) bool {
	// check for cm or in substring
	if strings.Contains(height, "in") {
		h := strings.Split(height, "in")
		heightInt, _ := strconv.Atoi(h[0])
		if heightInt >= 59 && heightInt <= 76 {
			return true
		}
	} else if strings.Contains(height, "cm") {
		h := strings.Split(height, "cm")
		heightInt, _ := strconv.Atoi(h[0])
		if heightInt >= 150 && heightInt <= 193 {
			return true
		}
	}

	return false
}

func validHair(hairColor string) bool {
	if len(hairColor) != 7 {
		return false
	}

	hc := strings.Split(hairColor, "")
	if hc[0] != "#" {
		return false
	}

	for i := 1; i < len(hc); i++ {
		if !isInt(hc[i]) {
			switch hc[i] {
			case "a":
				continue
			case "b":
				continue
			case "c":
				continue
			case "d":
				continue
			case "e":
				continue
			case "f":
				continue
			default:
				return false
			}
		}
	}

	return true
}

func validEye(eyeColor string) bool {
	switch eyeColor {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}
	return false
}

func validID(id string) bool {
	if len(id) != 9 {
		return false
	}
	split := strings.Split(id, "")
	for _, v := range split {
		if !isInt(v) {
			return false
		}
	}
	return true
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
