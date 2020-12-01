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