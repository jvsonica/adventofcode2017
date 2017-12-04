package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("day3")

	fmt.Println(Problem1(12), "should be 3")
	fmt.Println(Problem1(23), "should be 2")
	fmt.Println(Problem1(1024), "should be 31")
	fmt.Println("the problem result is", Problem1(312051))

}

func Problem1(value int) int{
	// Find the side size of the spiral
	dim := 1

	for dim*dim <= value {
		dim = dim + 2
	}

	// Depending on which side of the grid the value is we ite
	var x int
	var y int
	var currentValue int

	if value > (dim-1)*(dim-1) {
		x = dim
		y = dim
		currentValue = dim*dim

		for currentValue > value {
			// fmt.Println(x, y)
			if x == 1 {
				y--
			} else if y == dim {
				x--
			}
			currentValue--
		}
	} else {
		x = 2
		y = 1
		currentValue = (dim-1)*(dim-1)

		for currentValue > value {
			// fmt.Println(x, y)
			if x == dim {
				y++
			} else if y == 1 {
				x++
			}
			currentValue--
		}
	}

	// Center point of the grid (with the current dim)
	targetX := dim/2 + 1
	targetY := dim/2 + 1

	// Manhattan Distance to center
	distance := math.Abs(float64(x - targetX)) + math.Abs(float64(y - targetY))
	return int(distance)
}