package main

import (
	"fmt"
	"bufio"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	// Read file into string
	fileContent := ReadFile("input.txt")

	fmt.Println(CorruptionChecksum("5	1	9	5\n7	5	3\n2	4	6	8\n"), "should be 18")
	fmt.Println("the problem result is", CorruptionChecksum(fileContent))

	fmt.Println(CorruptionChecksum2("5	9	2	8\n9	4	7	3\n3	8	6	5	8\n"), "should be 9")
	fmt.Println("the problem result is", CorruptionChecksum2(fileContent))
}

func CorruptionChecksum(input string) int{
	// Difference of each line
	var differences []int

	// Run every line
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		temp := strings.Split(currentLine,"\t")
		min, max := minAndMax(stringArrToIntArray(temp))
		differences = append(differences, max-min)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return sum(differences)
}

func CorruptionChecksum2(input string) int{
	// Quotient of each line
	var divisions []int

	// Run every line
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		temp := strings.Split(currentLine,"\t")
		div1, div2 := findDivisible(stringArrToIntArray(temp))
		divisions = append(divisions, div1/div2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return sum(divisions)
}


func ReadFile(path string) string{
	// Read file
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	return string(fileContent)
}

func findDivisible(array []int) (int, int) {
	divisible1 := array[0]
	divisible2 := array[0]

	for _, value := range array {
		currentDivisible1 := value
		for _, currentDivisible2 := range array {
			if currentDivisible2 == currentDivisible1 {
				continue
			}
			if currentDivisible1 % currentDivisible2 == 0 {
				divisible1 = currentDivisible1
				divisible2 = currentDivisible2
			}
		}
	}

	return divisible1, divisible2
}

func minAndMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func sum(arr []int) int {
	// Calculate sum of int array
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func stringArrToIntArray(arr []string) []int{
	numberArray := make([] int, 0, len(arr))
	for _, a := range arr {
		i, _ := strconv.Atoi(string(a))
		numberArray = append(numberArray, i)
	}
	return numberArray
}
