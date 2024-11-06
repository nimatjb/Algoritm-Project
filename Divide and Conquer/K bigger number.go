package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSelect(arr []int, low, high, k int) int {
	if low <= high {
		pi := partition(arr, low, high)
		if pi == k-1 {
			return arr[pi]
		} else if pi > k-1 {
			return quickSelect(arr, low, pi-1, k)
		} else {
			return quickSelect(arr, pi+1, high, k)
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("enter your number:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	nums := strings.Fields(input)

	arr := make([]int, len(nums))
	for i, num := range nums {
		val, err := strconv.Atoi(num)
		if err != nil {
			//fmt.Println("your list not sort.")
			return
		}
		arr[i] = val
	}

	//fmt.Println("enter your k:")
	kInput, _ := reader.ReadString('\n')
	kInput = strings.TrimSpace(kInput)
	k, err := strconv.Atoi(kInput)
	if err != nil || k <= 0 || k > len(arr) {
		//fmt.Println("your k is not crect")
		return
	}

	result := quickSelect(arr, 0, len(arr)-1, k)
	fmt.Println(result)

}
