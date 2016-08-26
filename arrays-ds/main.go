package main

import (
	"fmt"
	"log"
)

func main() {
	length, err := readLength()
	if err != nil {
		log.Fatal(err)
	}
	nums, _ := readNums(length)
	if err != nil {
		log.Fatal(err)
	}
	printReverse(nums)
}

func readLength() (int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return -1, err
	}

	return length, nil
}

func readNums(length int) ([]int, error) {
	nums := make([]int, length)

	for i := range nums {
		_, err := fmt.Scanf("%d", &nums[i])
		if err != nil {
			return nil, err
		}
	}
	return nums, nil
}

func printReverse(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("%d", nums[i])
		if i != 0 {
			fmt.Printf(" ")
		}
	}
}
