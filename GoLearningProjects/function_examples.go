package main

import "math"

func addition(x int, y int) int {
	return x + y
}

func multiplication(x int, y int) (product int) {
	product = x * y
	return
}

func division(x int, y int) float32 {
	return float32(x / y)
}

func difference(x float64, y float64) float64 {
	return math.Abs(x - y)
}
