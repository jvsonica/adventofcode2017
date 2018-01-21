package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// Read test and input.txt files
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(fileContent)
	lengths := stringToIntArray(input)

	// Run
	fmt.Println("Part 1")
	fmt.Println(Problem1([]int{3,4,1,5}, makeRange(0, 4)), "should be 12")
	fmt.Println("the problem result is", Problem1(lengths, makeRange(0, 255)))

}


func Problem1(inputLengths []int, numberList []int) int {
	skipSize := 0
	currentPosition := 0

	for i := 0; i < len(inputLengths); i++ {
		length := inputLengths[i]

		for j := 0; j < length / 2; j++ {
			a := (currentPosition + j) % len(numberList)
			b := (currentPosition + length -1 - j) % len(numberList)
			numberList[a], numberList[b] = numberList[b], numberList[a]
		}

		currentPosition = (currentPosition + length + skipSize) % len(numberList)
		skipSize = skipSize + 1
	}

	return numberList[0] * numberList[1]
}

func stringToIntArray(input string) []int{
	numberArray := make([] int, 0)
	separatedStrings := strings.Split(input, ",")
	for _, str := range separatedStrings {
		value, _ := strconv.Atoi(str)
		numberArray = append(numberArray, value)
	}
	return numberArray
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}