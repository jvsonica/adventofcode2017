package main

import (
	"io/ioutil"
	"fmt"
	"bufio"
	"strings"
	"sort"
)

func main() {
	// Read file
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(fileContent)

	fmt.Println("Part 1")
	fmt.Println(Problem1("aa bb cc dd ee\n"), "should be 1")
	fmt.Println(Problem1("aa bb cc dd aa\n"), "should be 0")
	fmt.Println(Problem1("aa bb cc dd aaa\n"), "should be 1")
	fmt.Println("the problem result is", Problem1(input))

	fmt.Println("Part 2")
	fmt.Println(Problem2("abcde fghij\n"), "should be 1")
	fmt.Println(Problem2("abcde xyz ecdab\n"), "should be 0")
	fmt.Println(Problem2("a ab abc abd abf abj\n"), "should be 1")
	fmt.Println(Problem2("iiii oiii ooii oooi oooo\n"), "should be 1")
	fmt.Println(Problem2("oiii ioii iioi iiio\n"), "should be 0")
	fmt.Println("the problem result is", Problem2(input))
}

func Problem1(input string) int{
	counter := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()

		// Split each line's words into an array
		words := strings.Split(currentLine," ")

		if AreElementsUnique(words) {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return counter
}

func Problem2(input string) int{
	counter := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		currentLine := scanner.Text()
		words := strings.Split(currentLine," ")
		sortedWords := make([]string, len(words))

		for i, word := range words {
			characters := strings.Split(word, "")
			sort.Strings(characters)
			sortedWords[i] = strings.Join(characters, "")
		}

		if AreElementsUnique(sortedWords) {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Print(err)
	}

	return counter
}

func AreElementsUnique(arr []string) bool {
	// Create a map that will contain the each word as Key
	auxMap := make(map[string]bool)

	// Only non existent words will be added to the Map
	for i:=0; i < len(arr); i++ {
		auxMap[arr[i]] = true
	}

	// If the length matches than all words are unique
	return len(auxMap) == len(arr)
}