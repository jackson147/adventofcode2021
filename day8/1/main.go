package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func isSimpleDigit(signal string) bool {
	signalLength := len(signal)
	//1 4 7 8
	return signalLength == 2 || signalLength == 4 || signalLength == 3 || signalLength == 7
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day8/1/real.txt")
	check(err)

	simpleOutputSignalCount := 0
	for _, line := range lines {
		inputOutput := strings.Split(line, " | ")

		inputs := strings.Split(inputOutput[0], " ")
		outputs := strings.Split(inputOutput[1], " ")

		fmt.Println(inputs, outputs)

		for _, output := range outputs {
			if isSimpleDigit(output) {
				simpleOutputSignalCount++
			}
		}
	}
	fmt.Println("Single digit count: ", simpleOutputSignalCount)
}
