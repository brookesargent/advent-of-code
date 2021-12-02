package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	numbers, err := helper.ReadInputTxtToIntSlice("9/input.txt")
	if err != nil {
		log.Println(err)
	}

	preambleLength := 25
	preamble := numbers[0:preambleLength]

	weakness := 0
	for i := preambleLength; i < len(numbers); i++ {
		if !isSummed(numbers[i], preamble) {
			weakness = numbers[i]
			break
		}
		preamble = pop(preamble)
		preamble = append(preamble, numbers[i])
	}

	high, low := encryptionWeakness(numbers, weakness)
	fmt.Println("The part 1 weakness is: " + strconv.Itoa(weakness))
	fmt.Println("The part 2 weakness is: " + strconv.Itoa(high+low))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func encryptionWeakness(numbers []int, weakness int) (int, int) {
	var high, low int
	for i, v := range numbers {
		summation := v
		high = v
		low = v
		for n := 1; true; n++ {
			current := numbers[i+n]
			summation += current
			if high < current {
				high = current
			}
			if low > current {
				low = current
			}
			if summation == weakness {
				return high, low
			}
			if summation > weakness {
				break
			}
		}
	}
	return high, low
}

func isSummed(sum int, preamble []int) bool {
	seen := make([]int, 0)
	for _, num := range preamble {
		if contains(seen, sum-num) {
			return true
		}
		seen = append(seen, num)
	}
	return false
}

func contains(slice []int, i int) bool {
	for _, v := range slice {
		if v == i {
			return true
		}
	}
	return false
}

func pop(a []int) []int {
	_, a = a[0], a[1:]
	return a
}
