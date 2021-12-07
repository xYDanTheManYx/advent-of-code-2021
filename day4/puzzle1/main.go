package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoRowNumber struct {
	number int
	marked bool
}

type bingoRow []bingoRowNumber

type bingoBoard []bingoRow

func bingoNumbers(filename string) []int {
	var numbers []int
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
		numberList := strings.Split(scanner.Text(), ",")
		for _, number := range numberList {
			n, _ := strconv.Atoi(number)
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func bingoBoards(filename string) []bingoBoard {
	var bingoBoards []bingoBoard
	f, err := os.Open(filename)

	if err != nil {

	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var board bingoBoard
	var rows []bingoRow

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		var row []string
		rawRow := strings.Split(scanner.Text(), " ")
		for _, number := range rawRow {
			if number != "" {
				row = append(row, number)
			} else {
				continue
			}
		}
		if row != nil {
			number1, _ := strconv.Atoi(row[0])
			number2, _ := strconv.Atoi(row[1])
			number3, _ := strconv.Atoi(row[2])
			number4, _ := strconv.Atoi(row[3])
			number5, _ := strconv.Atoi(row[4])
			rows = append(rows, bingoRow{bingoRowNumber{number1, false}, bingoRowNumber{number2, false},
				bingoRowNumber{number3, false}, bingoRowNumber{number4, false}, bingoRowNumber{number5, false}})
		} else {
			board = bingoBoard{rows[0], rows[1], rows[2], rows[3], rows[4]}
			bingoBoards = append(bingoBoards, board)
			rows = []bingoRow{}
			continue
		}

	}
	return bingoBoards
}

func checkBingoCard(board bingoBoard) bool {
	for _, checkRow := range board {
		if checkRow[0].marked && checkRow[1].marked && checkRow[2].marked && checkRow[3].marked && checkRow[4].marked == true {
			return true
		}
		for i := 0; i < 5; i++ {
			if board[0][i].marked && board[1][i].marked && board[2][i].marked && board[3][i].marked && board[4][i].marked == true {
				return true
			}
		}
	}
	return false
}

func answer(board bingoBoard, bingoNumber int) int {
	var numbers []int
	var answer int
	for _, answerRow := range board {
		for _, answerNumber := range answerRow {
			if answerNumber.marked == false {
				numbers = append(numbers, answerNumber.number)
			}
		}
	}
	for _, unmarkedNumber := range numbers {
		answer += unmarkedNumber
	}
	return answer * bingoNumber
}

func main() {
	bingoNumbers := bingoNumbers("input_numbers.txt")
	//fmt.Println(bingoNumbers)
	bingoBoards := bingoBoards("input_boards.txt")
	//fmt.Println(bingoBoards)

	for _, bingoNumber := range bingoNumbers {
		for bingoBoardIndex, bingoBoard := range bingoBoards {
			for bingoCardIndex, bingoCardRow := range bingoBoard {
				for bingoCardRowIndex, bingoCardNumber := range bingoCardRow {
					if bingoCardNumber.number == bingoNumber {
						bingoBoards[bingoBoardIndex][bingoCardIndex][bingoCardRowIndex].marked = true
						if checkBingoCard(bingoBoard) {
							fmt.Println(bingoNumber)
							fmt.Println(bingoBoard)
							fmt.Println(answer(bingoBoard, bingoNumber))
							return
						}
					}
				}
			}
		}
	}
}
