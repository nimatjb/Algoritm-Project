package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print()
	input, _ := reader.ReadString('\n')

	strNums := strings.Fields(input)
	nums := make([]int, len(strNums))
	for i, str := range strNums {
		nums[i], _ = strconv.Atoi(str)
	}

	quickSort(nums, 0, len(nums)-1)

	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
