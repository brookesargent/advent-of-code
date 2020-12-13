package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/brookesargent/advent-of-code2020/helper"
)

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("13/input.txt")
	if err != nil {
		log.Println(err)
	}

	departureTimestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Printf("%s is not a valid number: %s\n", lines[0], err)
	}

	busIds := strings.Split(lines[1], ",")

	earliestDeparture := 0
	busToTake := 0
	for _, id := range busIds {
		currentID, _ := strconv.Atoi(id)
		d := float64(departureTimestamp) / float64(currentID)
		departure := int(math.Ceil(d)) * currentID
		if departure >= departureTimestamp && departure < earliestDeparture || earliestDeparture == 0 {
			earliestDeparture = departure
			busToTake = currentID
		}
	}
	fmt.Println((earliestDeparture - departureTimestamp) * busToTake)
	fmt.Println("Program duration: " + time.Since(start).String())
}
