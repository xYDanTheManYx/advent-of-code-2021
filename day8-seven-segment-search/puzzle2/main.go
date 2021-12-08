package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func allSegmentNumbers(filename string) [][][]string {
	var allSegmentNumbers [][][]string
	f, err := os.Open(filename)

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

		var segmentNumbers []string
		var fourDigit []string
		var afterDelimiter bool
		var row [][]string

		numbers := strings.Split(scanner.Text(), " ")
		for _, number := range numbers {
			if number == "|" {
				afterDelimiter = true
			} else {
				if afterDelimiter {
					fourDigit = append(fourDigit, number)
				} else {
					segmentNumbers = append(segmentNumbers, number)
				}
			}
		}
		row = append(row, segmentNumbers, fourDigit)
		allSegmentNumbers = append(allSegmentNumbers, row)
	}
	return allSegmentNumbers
}

func numberSegments(segments [10]string, remainingNumbers []string) [10]string {
	for _, remainingNumber := range remainingNumbers {
		if len(remainingNumber) == 6 { // 0, 6 or 9
			if sixCheck(remainingNumber, segments[1]) {
				segments[6] = remainingNumber
			} else if nineCheck(remainingNumber, segments[4]) {
				segments[9] = remainingNumber
			} else {
				segments[0] = remainingNumber
			}
		}
	}

	for _, remainingNumber := range remainingNumbers {
		if len(remainingNumber) == 5 { // 2, 3 or 5
			if twoCheck(remainingNumber, segments[9]) {
				segments[2] = remainingNumber
			} else if fiveCheck(remainingNumber, segments[6]) {
				segments[5] = remainingNumber
			} else {
				segments[3] = remainingNumber
			}
		}
	}

	return segments
}

func sixCheck(number, one string) bool {
	var count int
	for _, letter := range one {
		for _, checkLetter := range number {
			if letter == checkLetter {
				count += 1
			}
		}
	}
	if count != 2 {
		return true
	}
	return false
}

func nineCheck(number, four string) bool {
	var count int
	for _, letter := range four {
		for _, checkLetter := range number {
			if letter == checkLetter {
				count += 1
			}
		}
	}
	if count == 4 {
		return true
	}
	return false
}

func twoCheck(number, nine string) bool {
	var count int
	for _, letter := range number {
		for _, checkLetter := range nine {
			if letter == checkLetter {
				count += 1
			}
		}
	}
	if count != 5 {
		return true
	}
	return false
}

func fiveCheck(number, six string) bool {
	var count int
	for _, letter := range number {
		for _, checkLetter := range six {
			if letter == checkLetter {
				count += 1
			}
		}
	}
	if count == 5 {
		return true
	}
	return false
}

func findNumber(number string, numbers [10]string) int {
	var count int
	for index, numbersSegment := range numbers {
		if len(number) == len(numbersSegment) {
			count = 0
			for _, numberSegmentLetter := range number {
				for _, numbersSegmentLetter := range numbersSegment {
					if numberSegmentLetter == numbersSegmentLetter {
						count += 1
					}
				}
			}
		}
		if count == len(number) {
			return index
		}
	}
	panic("no number found")
}

func main() {
	allSegmentNumbers := allSegmentNumbers("input.txt")
	var allOutputNumbers []int

	for _, row := range allSegmentNumbers {

		var number [10]string
		var otherNumbers []string

		for _, segmentNumber := range row[0] {
			switch len(segmentNumber) {
			case 2:
				number[1] = segmentNumber
			case 4:
				number[4] = segmentNumber
			case 3:
				number[7] = segmentNumber
			case 7:
				number[8] = segmentNumber
			default:
				otherNumbers = append(otherNumbers, segmentNumber)
			}
		}
		numberSegments := numberSegments(number, otherNumbers)
		var outputNumber string

		for _, output := range row[1] {
			var foundNumber int
			foundNumber = findNumber(output, numberSegments)
			outputNumber += strconv.Itoa(foundNumber)
		}
		outputNumberInt, _ := (strconv.Atoi(outputNumber))
		allOutputNumbers = append(allOutputNumbers, outputNumberInt)
	}
	var outputNumberTotal int
	for _, outputNumber := range allOutputNumbers {
		outputNumberTotal += outputNumber
	}
	fmt.Println(outputNumberTotal)
}
