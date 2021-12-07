package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int
	var threeNums []int
	var increased int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a pattern to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	for i := range nums {
		if i == 1999 || i == 1998 {
			continue
		}
		threeNums = append(threeNums, nums[i]+nums[i+1]+nums[i+2])
	}

	for i, threeNum := range threeNums {
		if i == 0 {
			continue
		}
		if threeNum > threeNums[i-1] {
			increased += 1
		}
	}
	fmt.Println(increased)
}
