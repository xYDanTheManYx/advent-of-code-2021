package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction struct {
	direction string
	amount    int
}

func main() {
	var directions []direction

	f, err := os.Open("input.txt")

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

		s := strings.Split(scanner.Text(), " ")
		var inputDirection direction
		inputAmount, _ := strconv.Atoi(s[1])
		inputDirection.direction, inputDirection.amount = s[0], inputAmount
		directions = append(directions, inputDirection)

	}

	horizontal := 0
	vertical := 0

	for _, d := range directions {
		switch d.direction {
		case "up":
			vertical -= d.amount
		case "down":
			vertical += d.amount
		case "forward":
			horizontal += d.amount
		}
	}

	fmt.Println(horizontal * vertical)
}
