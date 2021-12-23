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

func copyMap(toCopy map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range toCopy {
		newMap[key] = value
	}
	return newMap
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day14/2/example.txt")
	// lines, err := readLines("/home/rich/git/advent2021/day14/2/real.txt")
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

	pairs := make([][]string, 0)
	currentPair := make([]string, 0)
	resultMap := make(map[string]int)
	for _, instructionRune := range instructionLine {
		currentPair = append(currentPair, string(instructionRune))
		if len(currentPair) == 2 {
			pairs = append(pairs, currentPair)
			currentPair = currentPair[1:]
		}
		resultMap[string(instructionRune)] = resultMap[string(instructionRune)] + 1
	}

	// fmt.Println(resultMap)
	costMap := make(map[string]int)
	for i, pair := range pairs {
		fmt.Println(i)
		processPairs(polymerMap, resultMap, pair, 0, costMap)
	}
	fmt.Println(resultMap)
	values := make([]int, 0)
	for _, value := range resultMap {
		values = append(values, value)
	}
	// fmt.Println(values)
	min, max := minMax(values)
	fmt.Println(max - min)
}

func processPairs(polymerMap map[string]string, resultMap map[string]int, pair []string, iteration int, costMap map[string]int) {
	combined := pair[0] + pair[1]
	newValue := polymerMap[combined]
	resultMap[newValue] = resultMap[newValue] + 1

	costMap[newValue] = costMap[newValue] + 1

	//Create new pairs
	var newPairOne = []string{pair[0], newValue}
	var newPairTwo = []string{newValue, pair[1]}
	iteration++

	if iteration == 40 {
		return
	}

	//Send for more map entries..
	processPairs(polymerMap, resultMap, newPairOne, iteration, copyMap(costMap))
	processPairs(polymerMap, resultMap, newPairTwo, iteration, copyMap(costMap))
}
