package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]

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

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	// lines, err := readLines("/home/rich/git/advent2021/day14/1/example.txt")
	lines, err := readLines("/home/rich/git/advent2021/day14/1/real.txt")
	check(err)

	instructionLine := ""

	polymerMap := make(map[string]string)

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			instructionLine = line
		} else {
			polymers := strings.Split(line, " -> ")
			polymerMap[polymers[0]] = polymers[1]
		}
	}

	fmt.Println(instructionLine)
	fmt.Println(polymerMap)

	for i := 0; i < 10; i++ {
		instructionLine = getResult(polymerMap, instructionLine)
	}
	// fmt.Println(instructionLine)
	// fmt.Println(instructionLine == "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")

	resultMap := calculateRuneOccurances(instructionLine)
	// fmt.Println(resultMap)
	values := make([]int, 0)
	for _, value := range resultMap {
		values = append(values, value)
	}
	// fmt.Println(values)
	min, max := minMax(values)
	fmt.Println(max - min)
}

func getResult(polymerMap map[string]string, instructionLine string) string {
	polyString := ""
	outputString := ""
	for i, input := range instructionLine {
		polyString += string(input)
		if len(polyString) == 2 {
			result := polymerMap[polyString]
			outputString += string(polyString[0]) + result
			polyString = polyString[1:]
		}

		if i == len(instructionLine)-1 {
			outputString += string(input)
		}
	}
	return outputString
}

func calculateRuneOccurances(input string) map[rune]int {
	resultMap := make(map[rune]int)
	for _, stringRune := range input {
		resultMap[stringRune] = resultMap[stringRune] + 1
	}
	return resultMap
}
