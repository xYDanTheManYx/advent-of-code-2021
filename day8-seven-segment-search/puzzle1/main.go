package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func allSegmentNumbers(filename string) [][][]string {
	var allSegmentNumbers [][][]string
	f, err := os.Open(filename)

	if err != nil {

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

func main() {
	allSegmentNumbers := allSegmentNumbers("input.txt")

	var counter int

	for _, row := range allSegmentNumbers {
		for _, number := range row[1] {
			if len(number) == 2 || len(number) == 4 || len(number) == 3 || len(number) == 7 {
				counter += 1
			}
		}
	}
	fmt.Println(counter)
}
