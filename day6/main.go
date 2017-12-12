package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// Read file
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(fileContent)
	// fmt.Println(input)

	fmt.Println("Part 1")
	fmt.Println(Problem1("0\t2\t7\t0"), "should be 5")
	fmt.Println("the problem result is", Problem1(input))

	fmt.Println("Part 2")
	fmt.Println(Problem1("2\t4\t1\t2")-1, "should be 4")
	input = "14\t13\t12\t11\t9\t8\t8\t6\t6\t4\t4\t3\t1\t1\t0\t12"
	fmt.Println("the problem result is", Problem1(input)-1)
}

func Problem1(input string) int{
	// Redistribute it through other positions
	// Store State
	// Increment counter
	arr := StringToIntArray(input)
	var solutions [][]int
	counter := 0

	for AreElementsUnique(solutions) != false {
		// Pick largest value and its index
		position, max := Max(arr)

		// Get iteration order for current largest value
		var iterations []int
		iterations = append(iterations, makeRange(position+1, len(arr)-1)...)
		iterations = append(iterations, makeRange(0, position)...)

		// Empty the max value's position
		arr[position] = 0

		// Redistribute through all slots
		for i := 0; i < max; i++ {
			index := iterations[i % len(iterations)]
			arr[index] += 1
		}

		// Increment counter
		counter++

		// Save current configuration to check for the first repeated solution
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		solutions = append(solutions, tmp)
	}

	fmt.Println("last state:", arr)
	return counter
}

// Transform \t separated string of numbers into int array
func StringToIntArray(str string) []int {
	array := make([]int, 0)
	result := strings.Split(str, "\t")

	for i := range result {
		value, _ := strconv.Atoi(result[i])
		array = append(array, value)
	}


	return array
}

func Max(arr []int) (int, int) {
	max := arr[0]
	index := 0

	for i, value := range arr {
		if value > max {
			max = value
			index = i
		}
	}
	return index, max
}

func AreElementsUnique(arr [][]int) bool {
	// Create a map that will contain the each array as Key
	auxMap := make(map[string]bool)

	// Only non existent words will be added to the Map
	for i:=0; i < len(arr); i++ {
		key := fmt.Sprint(arr[i])
		auxMap[key] = true
	}

	// If the length matches than all words are unique
	return len(auxMap) == len(arr)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}