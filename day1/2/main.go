package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func sumArray(values []int) int{
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func createRollingWindowSlice(lines []string, windowSize int) []int {
	rollingWindow := make([]int, 0)
	outputValues := make([]int, 0)

	for i, line := range lines {

		value, err := strconv.Atoi(line)
		check(err)

		rollingWindow = append(rollingWindow, value)

		if len(rollingWindow) == windowSize {

			fmt.Println(i, rollingWindow)

			//Tot up and add to result slice
			sum := sumArray(rollingWindow)
			outputValues = append(outputValues, sum)

			// Discard top element
			rollingWindow = rollingWindow[1:]

			////Last loop, won't loop round again, so need to add this manually
			if i == len(lines){
				sum = sumArray(rollingWindow)
				outputValues = append(outputValues, sum)
				rollingWindow = append(rollingWindow, value)
			}
		}
		////Has to equal window size again
	}
	return outputValues
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day1/2/real.txt")
	check(err)

	rollingWindowSlice := createRollingWindowSlice(lines, 3)

	last := 0
	count := 0
	for i, current := range rollingWindowSlice {
		fmt.Println(i, current)

		if i > 0 && current > last {
			count++
		}

		last = current
	}
	fmt.Println(count)
}