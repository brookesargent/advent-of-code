package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInputTxtToIntSlice(filepath string) ([]int, error) {
	var lines []int
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		lines = append(lines, num)
	}
	return lines, nil
}

func ReadInputTxtToStringSlice(filepath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lines = append(lines, lineStr)
	}
	return lines, nil
}

func RemoveDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
