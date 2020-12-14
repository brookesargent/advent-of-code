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

type Bus struct {
	ID     int
	Offset int
}

func main() {
	start := time.Now()

	lines, err := helper.ReadInputTxtToStringSlice("13/input.txt")
	if err != nil {
		log.Println(err)
	}

	// part one
	departureTimestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Printf("%s is not a valid number: %s\n", lines[0], err)
	}
	earliestDeparture, busToTake := findEarliestDeparture(departureTimestamp, strings.Split(lines[1], ","))
	d := (earliestDeparture - departureTimestamp) * busToTake
	// part two
	buses := formatBusIds(lines[1])
	t := findT(buses)
	fmt.Println("Part one: " + strconv.Itoa(d))
	fmt.Println("Part two: " + strconv.Itoa(t))
	fmt.Println("Program duration: " + time.Since(start).String())
}

func findEarliestDeparture(timestamp int, busIds []string) (int, int) {
	earliestDeparture := 0
	busToTake := 0
	for _, id := range busIds {
		currentID, _ := strconv.Atoi(id)
		d := float64(timestamp) / float64(currentID)
		departure := int(math.Ceil(d)) * currentID
		if departure >= timestamp && departure < earliestDeparture || earliestDeparture == 0 {
			earliestDeparture = departure
			busToTake = currentID
		}
	}
	return earliestDeparture, busToTake
}

func findT(buses []Bus) int {
	t := 0
	busCount := 2
	offset := buses[0].ID
	// works to find starting t - how do i make it work for 3, 4, 5 etc buses?
	for busCount <= len(buses) {
	count:
		for i := t; true; i += offset {
			for j := 0; j < busCount; j++ {
				if (i+buses[j].Offset)%buses[j].ID != 0 {
					break
				} else if j == busCount-1 {
					t = i
					break count
				}
			}
		}
		offset = getOffset(busCount, buses)
		busCount++
	}
	return t
}

func getOffset(busCount int, buses []Bus) int {
	offset := 1
	for i := 0; i < busCount; i++ {
		offset *= buses[i].ID
	}
	return offset
}

func formatBusIds(line string) []Bus {
	split := strings.Split(line, ",")
	buses := make([]Bus, 0)
	for i, v := range split {
		id, err := strconv.Atoi(v)
		if err == nil {
			buses = append(buses, Bus{
				ID:     id,
				Offset: i,
			})
		}
	}
	return buses
}
