package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func crabSubmarines(filename string) []int {
	var numbers []int
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
		numberList := strings.Split(scanner.Text(), ",")
		for _, number := range numberList {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func biggestNumber(array []int) int {
	var biggestNumber int
	for _, number := range array {
		if number > biggestNumber {
			biggestNumber = number
		}
	}
	return biggestNumber
}

func main() {
	crabSubmarines := crabSubmarines("input.txt")

	biggestNumber := biggestNumber(crabSubmarines)

	var horizonalLevel int
	var fuel int

	for i := 0; i < biggestNumber+1; i++ {
		fuelNeeded := 0
		for _, crabSubmarine := range crabSubmarines {
			up := 0
			down := 0
			if crabSubmarine < i {
				up = i - crabSubmarine
				fuelNeeded += up
			} else {
				down = crabSubmarine - i
				fuelNeeded += down
			}
		}
		if i == 0 {
			fuel = fuelNeeded
		} else {
			if fuelNeeded < fuel {
				fuel = fuelNeeded
				horizonalLevel = i
			}
		}
	}
	fmt.Println("Horizontal Level:", horizonalLevel, "Fuel", fuel)
}
