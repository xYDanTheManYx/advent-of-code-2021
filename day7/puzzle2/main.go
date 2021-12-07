package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var crabSubmarinesfloat64 []float64

	for _, crabSubmarine := range crabSubmarines {
		crabSubmarinesfloat64 = append(crabSubmarinesfloat64, float64(crabSubmarine))
	}

	var horizonalLevel int
	var fuel float64

	for i := 0; i < biggestNumber+1; i++ {
		fuelNeeded := 0.0
		for _, crabSubmarine := range crabSubmarinesfloat64 {
			up := 0.0
			down := 0.0
			if crabSubmarine < float64(i) {
				up = float64(i) - crabSubmarine
				fuelNeeded += (math.Pow(up, 2) + up) / 2
			} else {
				down = crabSubmarine - float64(i)
				fuelNeeded += (math.Pow(down, 2) + down) / 2
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
	finalFuel := int(fuel)
	fmt.Println("Horizontal Level:", horizonalLevel, "Fuel", finalFuel)
}
