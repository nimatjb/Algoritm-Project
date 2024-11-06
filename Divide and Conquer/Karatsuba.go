package main

import (
	"fmt"
	"math"
	"strconv"
)

func karatsuba(a, b int) int {
	if a < 10 || b < 10 {
		return a * b
	}

	n := int(math.Max(float64(len(strconv.Itoa(a))), float64(len(strconv.Itoa(b)))))
	mid := n / 2

	highA := a / int(math.Pow(10, float64(mid)))
	lowA := a % int(math.Pow(10, float64(mid)))
	highB := b / int(math.Pow(10, float64(mid)))
	lowB := b % int(math.Pow(10, float64(mid)))

	z0 := karatsuba(lowA, lowB)
	z2 := karatsuba(highA, highB)
	z1 := karatsuba(lowA+highA, lowB+highB) - z2 - z0

	return (z2 * int(math.Pow(10, float64(2*mid)))) + (z1 * int(math.Pow(10, float64(mid)))) + z0
}

func main() {
	var x, y int
	fmt.Print()
	fmt.Scan(&x)
	fmt.Print()
	fmt.Scan(&y)

	fmt.Println(karatsuba(x, y))
}
