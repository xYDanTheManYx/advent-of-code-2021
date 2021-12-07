package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func lanternFish(filename string) []int {
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

func lanternFishDay(totalLanternFish [9]int) [9]int {
	var lanternFishEndOfDay [9]int
	for i := len(totalLanternFish) - 1; i >= 0; i-- {
		if i == 0 {
			lanternFishEndOfDay[8] = totalLanternFish[0]
			lanternFishEndOfDay[6] += totalLanternFish[0]
		} else {
			lanternFishEndOfDay[i-1] = totalLanternFish[i]
		}
	}
	return lanternFishEndOfDay
}

func main() {
	var totalLanternFish [9]int
	lanternFishes := lanternFish("input.txt")

	for _, lanternfish := range lanternFishes {
		totalLanternFish[lanternfish] += 1
	}

	for i := 0; i < 256; i++ {
		totalLanternFish = lanternFishDay(totalLanternFish)
	}

	var numberOfFish int
	for _, numberOfLanternFish := range totalLanternFish {
		numberOfFish += numberOfLanternFish
	}
	fmt.Println(numberOfFish)
}
