package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func coordinatesInput(filename string) [][]coordinate {
	var coordinateList [][]coordinate
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
		leftRight := strings.Split(scanner.Text(), " -> ")
		left := strings.Split(leftRight[0], ",")
		right := strings.Split(leftRight[1], ",")
		leftx, _ := strconv.Atoi(left[0])
		lefty, _ := strconv.Atoi(left[1])
		rightx, _ := strconv.Atoi(right[0])
		righty, _ := strconv.Atoi(right[1])
		coordinates := []coordinate{{leftx, lefty}, {rightx, righty}}
		coordinateList = append(coordinateList, coordinates)
	}
	return coordinateList
}

func verticalHorizontalLines(coordinates [][]coordinate) [][]coordinate {
	var coordinateList [][]coordinate
	for _, coordinate := range coordinates {
		if coordinate[0].x == coordinate[1].x || coordinate[0].y == coordinate[1].y {
			coordinateList = append(coordinateList, coordinate)
		}
	}
	return coordinateList
}

func highestXhighestY(coordinates [][]coordinate) (int, int) {
	x := 0
	y := 0
	for _, coordinate := range coordinates {
		if coordinate[0].x > x {
			x = coordinate[0].x
		}
		if coordinate[1].x > x {
			x = coordinate[1].x
		}
		if coordinate[0].y > y {
			y = coordinate[0].y
		}
		if coordinate[1].y > y {
			y = coordinate[1].y
		}
	}
	return x, y
}

func lowestHighestValue(number1, number2 int) (int, int) {
	if number1 <= number2 {
		return number1, number2
	} else {
		return number2, number1
	}
}

func main() {
	input := coordinatesInput("input.txt")
	verticalHorizontalLines := verticalHorizontalLines(input)
	fmt.Println(highestXhighestY(verticalHorizontalLines))

	var grid [991][991]int

	for _, verticalHorizontalLine := range verticalHorizontalLines {
		if verticalHorizontalLine[0].x == verticalHorizontalLine[1].x {
			startY, endY := lowestHighestValue(verticalHorizontalLine[0].y, verticalHorizontalLine[1].y)
			for i := startY; i < (endY + 1); i++ {
				grid[verticalHorizontalLine[0].x][i] += 1
			}
		} else {
			startX, endX := lowestHighestValue(verticalHorizontalLine[0].x, verticalHorizontalLine[1].x)
			for i := startX; i < (endX + 1); i++ {
				grid[i][verticalHorizontalLine[0].y] += 1
			}
		}
	}

	var points int
	for _, rows := range grid {
		for _, row := range rows {
			if row > 1 {
				points += 1
			}
		}
	}
	fmt.Println(points)
}
