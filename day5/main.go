package main

import (
	"io/ioutil"
	"fmt"
	"bufio"
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

	fmt.Println("Part 1")
	fmt.Println(Problem1("0\n3\n0\n1\n-3"), "should be 5")
	fmt.Println("the problem result is", Problem1(input))

	fmt.Println("Part 2")
	fmt.Println(Problem2("0\n3\n0\n1\n-3"), "should be 10")
	fmt.Println("the problem result is", Problem2(input))
}

func Problem1(input string) int{
	// Array with array that will be scanned
	array := StringToIntArray(input)

	counter := 0
	currentIt := 0
	solutionFound := false

	// Loop until we find a solution
	for solutionFound == false {
		counter++

		// Calculate next it
		nextIt := array[currentIt] + currentIt

		if nextIt > len(array) - 1 || nextIt < 0 {
			// If the next it is larger than the array scope
			// then we found the exit
			solutionFound = true
		} else {
			// Otherwise we keep going with the new it and change
			// the current position
			array[currentIt] = array[currentIt] + 1
			currentIt = nextIt
		}
	}

	return counter
}

func Problem2(input string) int{
	// Array with array that will be scanned
	array := StringToIntArray(input)

	counter := 0
	currentIt := 0
	solutionFound := false

	// Loop until we find a solution
	for solutionFound == false {
		counter++

		// Calculate next it
		nextIt := array[currentIt] + currentIt

		if nextIt > len(array) - 1 || nextIt < 0 {
			// If the next it is larger than the array scope
			// then we found the exit
			solutionFound = true
		} else {
			// Otherwise we keep going with the new it and change
			// the current position. For Problem2 the change is
			// a little different
			if array[currentIt] >= 3 {
				array[currentIt] = array[currentIt] - 1
			} else {
				array[currentIt] = array[currentIt] + 1
			}
			currentIt = nextIt
		}
	}

	return counter
}

// Transform \n separated string of numbers into int array
func StringToIntArray(str string) []int {
	array := make([]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		array = append(array, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return array
}
