package main

import (
    "fmt"
    "strconv"
    "io/ioutil"
)

func main() {
    // Read file
    fileContent, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Print(err)
    }

    input := string(fileContent)

    // Puzzle 1
    fmt.Println(SolveCaptcha("1122"), "should be 3")
    fmt.Println(SolveCaptcha("1111"), "should be 4")
    fmt.Println(SolveCaptcha("1234"), "should be 0")
    fmt.Println(SolveCaptcha("91212129"), "should be 9")
    fmt.Println("problem result is ", SolveCaptcha(input))

    // Puzzle 2
    fmt.Println(SolveCaptcha2("1212"), "should be 6")
    fmt.Println(SolveCaptcha2("1221"), "should be 0")
    fmt.Println(SolveCaptcha2("123425"), "should be 4")
    fmt.Println(SolveCaptcha2("123123"), "should be 12")
    fmt.Println(SolveCaptcha2("123131415"), "should be 4")
    fmt.Println("problem result is ", SolveCaptcha2(input))
}

func SolveCaptcha(input string) int{
    arr := stringToIntArray(input)
    arr = append(arr, arr[0])
    var consecutive []int

    for i:= 0; i < len(arr) - 1; i++ {
        if arr[i] == arr[i+1] {
            consecutive = append(consecutive, arr[i])
        }
    }

    return sum(consecutive)
}

func SolveCaptcha2(input string) int{
    arr := stringToIntArray(input)
    var halfwayConsecutive []int
    stepAhead := len(arr) / 2

    for i:=0; i < len(arr) / 2; i++ {
        if arr[i] == arr[stepAhead+i] {
            halfwayConsecutive = append(halfwayConsecutive, arr[i])
        }
    }

    return sum(halfwayConsecutive) * 2

}

func sum(arr []int) int {
    // Calculate sum of int array
    sum := 0
    for i := range arr {
        sum += arr[i]
    }
    return sum
}


func stringToIntArray(str string) []int{
    // Convert int's in a string to a int array 
    numberArray := make([] int, 0, len(str))
    for _, a := range str {
        i, _ := strconv.Atoi(string(a))
        numberArray = append(numberArray, i)
    }
    return numberArray
}
