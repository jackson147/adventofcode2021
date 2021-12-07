package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func getTotalFuelUsage(horizontalPosition int, crabPositions []int) int {
	fuelUsage := 0
	for _, crabPosition := range crabPositions {
		if crabPosition > horizontalPosition {
			fuelUsage += crabPosition - horizontalPosition
		} else {
			fuelUsage += horizontalPosition - crabPosition
		}
	}
	return fuelUsage
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day7/1/real.txt")
	check(err)

	crabPositions := make([]int, 0)
	for _, line := range lines {
		splitLine := strings.Split(line, ",")
		for _, stringValue := range splitLine {
			value, err := strconv.Atoi(stringValue)
			check(err)
			crabPositions = append(crabPositions, value)
		}
	}

	//Calculate the max and min values
	_, max := minMax(crabPositions)

	//Set the loop value double the size just in case
	loopMax := max * 2

	minFuelPosition := -1
	minFuelUsage := 999999999
	for i := 0; i < loopMax; i++ {
		fuelUsage := getTotalFuelUsage(i, crabPositions)
		if fuelUsage < minFuelUsage {
			minFuelUsage = fuelUsage
			minFuelPosition = i
		}
	}

	fmt.Println(minFuelPosition, minFuelUsage)
}
