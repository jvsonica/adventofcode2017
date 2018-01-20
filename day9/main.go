package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	// Read test and input files
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input := string(fileContent)

	// Running
	fmt.Println("Part 1")
	fmt.Println(Problem1("{}"), "should be 1")
	fmt.Println(Problem1("{{{}}}"), "should be 3")
	fmt.Println(Problem1("{{},{}}"), "should be 6")
	fmt.Println(Problem1("{{{},{},{{}}}}"), "should be 5")
	fmt.Println(Problem1("{<{},{},{{}}>}"), "should be 16")
	fmt.Println(Problem1("{<a>,<a>,<a>,<a>}"), "should be 1")
	fmt.Println(Problem1("{{<a>},{<a>},{<a>},{<a>}}"), "should be 9")
	fmt.Println(Problem1("{{<!>},{<!>},{<!>},{<a>}}"), "should be 3")
	fmt.Println("the problem result is", Problem1(input))
}

func Problem1(input string) int{
	scoreInsideGroup := 0
	score := 0

	for i := 0; i < len(input); i++ {
		currentChar := string(input[i])
		if currentChar ==  "{" {
			scoreInsideGroup = scoreInsideGroup + 1
			score = score + scoreInsideGroup
		} else if currentChar == "}" {
			scoreInsideGroup = scoreInsideGroup - 1
		} else if currentChar == "!" {
			i = i + 1
		} else if currentChar == "<" {
			endOfGarbage := i
			for j := i; j < len(input); j++ {
				if input[j] == '!' {
					j = j + 1
				} else if input[j] == '>' {
					endOfGarbage = j
					break
				}
			}
			i = endOfGarbage
		}
	}
	return score
}
