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

	zeroBitCount := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, number := range input {
		for bitPosition, bit := range number {
			if bit == 48 {
				zeroBitCount[bitPosition] += 1
			}
		}
	}

	gammaRate := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for index, zeroAmount := range zeroBitCount {
		if zeroAmount < 500 {
			gammaRate[index] = 1
		}
	}

	epsilonRate := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for index, bit := range gammaRate {
		if bit == 0 {
			epsilonRate[index] = 1
		}
	}

	gammaRateString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(gammaRate)), ""), "[]")
	episilonRateString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(epsilonRate)), ""), "[]")
	gammaRateDecimal, err := strconv.ParseInt(gammaRateString, 2, 64)
	epislonRateDecimal, err := strconv.ParseInt(episilonRateString, 2, 64)
	fmt.Println(gammaRateDecimal * epislonRateDecimal)
}
