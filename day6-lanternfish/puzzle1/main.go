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

func lanternFishDay(lanternFishes []int) []int {
	var newLanternFish int
	for index := range lanternFishes {
		if lanternFishes[index] == 0 {
			lanternFishes[index] = 7
			newLanternFish += 1
		}
		lanternFishes[index] = lanternFishes[index] - 1
	}
	for i := 0; i < newLanternFish; i++ {
		lanternFishes = append(lanternFishes, 8)
	}
	return lanternFishes
}

func main() {
	lanternFishes := lanternFish("input.txt")
	for i := 0; i < 80; i++ {
		lanternFishes = lanternFishDay(lanternFishes)
	}
	fmt.Println(len(lanternFishes))
}
