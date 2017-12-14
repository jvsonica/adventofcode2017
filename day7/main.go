package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
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
	fmt.Println(Problem1(inputTest), "should be tknk")
	fmt.Println("the problem result is", Problem1(input))

	fmt.Println("Part 2")
	fmt.Println(Problem2(inputTest), "should be 60")
	fmt.Println("the problem result is", Problem2(input))
}

func Problem1(input string) string{
	nodes := strings.Split(input, "\n")
	tree := createGraph(nodes)

	// Root is the node without parent
	var root string
	for _, node := range tree {
		if node.parent == nil {
			root = node.label
			break
		}
	}

	return root
}

func Problem2(input string) int{
	nodes := strings.Split(input, "\n")
	tree := createGraph(nodes)
	var result int

	for _, node := range tree {
		node.totalWeight = calculateWeight(node)
	}

	for _, node := range tree {
		if len(node.children) > 0 {
			// w := node.children[0].totalWeight
			weights := make(map[int]int)

			for _, child := range node.children {
				weights[child.totalWeight] = weights[child.totalWeight] + 1
			}

			if len(weights) > 1 {
				mostCommonWeight := findLargestValue(weights)
				for _, child := range node.children {
					if child.totalWeight != mostCommonWeight {
						incorrectNode := child
						result = incorrectNode.weight + mostCommonWeight - incorrectNode.totalWeight
					}
				}
			}
		}
	}

	return result
}

func findLargestValue(input map[int]int) int{
	largestValue := 0
	var result int
	for key, value := range input {
		if largestValue < value {
			result = key
		}
	}
	return result
}

func calculateWeight(node *node) int{
	if len(node.children) == 0 {
		return node.weight
	} else {
		totalWeight := node.weight
		for _, node := range node.children {
			totalWeight = totalWeight + calculateWeight(node)
		}
		return totalWeight
	}
}

func createGraph(nodeStrings []string) map[string]*node{
	tree := make(map[string]*node)

	for _, currentNode := range nodeStrings {
		// Splitting line
		nodeParts := strings.Split(currentNode, " ")
		name := nodeParts[0]
		weightString := nodeParts[1]
		weight, _ := strconv.Atoi(weightString[1 : len(weightString)-1])

		program := node{weight:weight, label:name}

		// If it has children, include in the program
		var children []string
		if strings.Contains(currentNode, "->") {
			children = strings.Split(currentNode, " -> ")
			children = strings.Split(children[1], ", ")

			program.childrenStr = children
		}

		// Append to program tree
		tree[name] = &program
	}

	for _, node := range tree {
		for _, nodeName := range node.childrenStr {
			// Assigning children
			node.children = append(node.children, tree[nodeName])

			// Assigning parent
			tree[nodeName].parent = node
		}
	}

	return tree
}

type node struct {
	label       string
	weight      int
	totalWeight int
	childrenStr []string
	children    []*node
	parent      *node
}

