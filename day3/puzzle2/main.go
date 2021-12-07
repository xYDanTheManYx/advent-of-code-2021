package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		binary := scanner.Text()
		input = append(input, binary)
	}

	oxygenNumbers := input
	oxygenZeroBitCount := zeroBitCounting(oxygenNumbers)
	var oxygenString string
	for i := 0; i < 12; i++ {
		oxygenNumbers = keepMostCommonBit(oxygenNumbers, oxygenZeroBitCount[i], i)
		if len(oxygenNumbers) == 1 {
			oxygenString = oxygenNumbers[0]
			break
		}
		oxygenZeroBitCount = zeroBitCounting(oxygenNumbers)
	}

	co2Numbers := input
	co2ZeroBitCount := zeroBitCounting(co2Numbers)
	var co2String string
	for i := 0; i < 12; i++ {
		co2Numbers = keepLeastCommonBit(co2Numbers, co2ZeroBitCount[i], i)
		if len(co2Numbers) == 1 {
			co2String = co2Numbers[0]
			break
		}
		co2ZeroBitCount = zeroBitCounting(co2Numbers)
	}

	oxygenGeneratorRating := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(oxygenString)), ""), "[]")
	co2ScrubberRating := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(co2String)), ""), "[]")
	oxygenGeneratorRatingDecimal, err := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	co2ScrubberRatingDecimal, err := strconv.ParseInt(co2ScrubberRating, 2, 64)
	fmt.Println(oxygenGeneratorRatingDecimal * co2ScrubberRatingDecimal)
}

func zeroBitCounting(input []string) []int {
	zeroBitCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, number := range input {
		for bitPosition, bit := range number {
			if bit == 48 {
				zeroBitCount[bitPosition] += 1
			}
		}
	}
	return zeroBitCount
}

func keepMostCommonBit(numbers []string, zeroBitCount int, bitCountPosition int) []string {
	var keep []string
	if zeroBitCount <= len(numbers)/2 {
		for _, number := range numbers {
			if number[bitCountPosition] == 49 {
				keep = append(keep, number)
			}
		}
	} else if zeroBitCount > len(numbers)/2 {
		for _, number := range numbers {
			if number[bitCountPosition] == 48 {
				keep = append(keep, number)
			}
		}
	}
	return keep
}

func keepLeastCommonBit(numbers []string, zeroBitCount int, bitCountPosition int) []string {
	var keep []string
	if zeroBitCount <= len(numbers)/2 {
		for _, number := range numbers {
			if number[bitCountPosition] == 48 {
				keep = append(keep, number)
			}
		}
	} else if zeroBitCount > len(numbers)/2 {
		for _, number := range numbers {
			if number[bitCountPosition] == 49 {
				keep = append(keep, number)
			}
		}
	}
	return keep
}
