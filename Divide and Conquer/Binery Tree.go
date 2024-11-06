package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binarySearch(arr []int, target int) string {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return strconv.Itoa(target)
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return "Not Found"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbersStr := scanner.Text()
	numbersSlice := strings.Fields(numbersStr)
	var numbers []int
	for _, numStr := range numbersSlice {
		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	result := binarySearch(numbers, target)
	fmt.Println(result)
}
