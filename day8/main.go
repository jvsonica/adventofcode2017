package main

import (
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"github.com/Knetic/govaluate"
)

func main() {
	// Read test and input files
	testContent, err := ioutil.ReadFile("input_test.txt")
	if err != nil {
		fmt.Println(err)
	}
	inputTest := string(testContent)

	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(fileContent)

	// Running
	fmt.Println("Part 1")
	fmt.Println(Problem1(inputTest), "should be 1")
	fmt.Println("the problem result is", Problem1(input))

	fmt.Println("Part 2")
	fmt.Println(Problem2(inputTest), "should be 10")
	fmt.Println("the problem result is", Problem2(input))
}

func Problem1(input string) int{
	instructions := readInstructions(input)
	parameters := make(map[string]interface{}, 0)

	// Initializing all variables
	for _, inst := range instructions {
		parameters[inst.variable] = 0
	}

	// Looping through instructions
	for _, inst := range instructions {
		expr, _ := govaluate.NewEvaluableExpression(inst.condition)
		result, _ := expr.Evaluate(parameters)

		// If instruction evaluates to true, then apply operation
		if result == true {
			if inst.operation == "inc" {
				parameters[inst.variable] = parameters[inst.variable].(int) + inst.amount
			} else if inst.operation == "dec" {
				parameters[inst.variable] = parameters[inst.variable].(int) - inst.amount
			}
		}
	}

	// Find max
	max := 0
	for _, value := range parameters {
		if value.(int) > max {
			max = value.(int)
		}
	}
	return max
}

func Problem2(input string) int{
	instructions := readInstructions(input)
	max := 0
	parameters := make(map[string]interface{}, 0)

	// Initializing all variables
	for _, inst := range instructions {
		parameters[inst.variable] = 0
	}

	// Looping through instructions
	for _, inst := range instructions {
		expr, _ := govaluate.NewEvaluableExpression(inst.condition)
		result, _ := expr.Evaluate(parameters)

		// If instruction evaluates to true, then apply operation
		if result == true {
			if inst.operation == "inc" {
				parameters[inst.variable] = parameters[inst.variable].(int) + inst.amount

			} else if inst.operation == "dec" {
				parameters[inst.variable] = parameters[inst.variable].(int) - inst.amount
			}

			// Saving max
			if parameters[inst.variable].(int) > max {
				max = parameters[inst.variable].(int)
			}
		}
	}

	return max
}

func readInstructions(input string) []Instruction{
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, 0)
	r, _ := regexp.Compile(`([^\s]+)\s([^\s]+)*\s(-?\d*)\sif\s(.*$)`)

	for _, line := range lines {
		// b inc 5 if a > 1
		groups := r.FindStringSubmatch(line)
		value, _ := strconv.Atoi(groups[3])

		currentInstruction := Instruction{variable:groups[1], operation:groups[2], condition:groups[4], amount:value}
		instructions = append(instructions, currentInstruction)
	}

	return instructions
}

type Instruction struct{
	condition string
	variable string
	operation string
	amount int
}