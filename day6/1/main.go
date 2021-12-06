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

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day6/1/real.txt")
	check(err)

	fishCycles := make([]int, 9)
	for _, line := range lines {

		splitLine := strings.Split(line, ",")
		for _, stringValue := range splitLine {
			var value, err = strconv.Atoi(stringValue)
			check(err)

			//Add fish to that cycle
			fishCycles[value] = fishCycles[value] + 1
		}
	}

	//Calculate fish in each lifecycle after number of days
	numberOfDays := 256
	for i := 0; i < numberOfDays; i++ {
		newFishCycles := make([]int, 9)

		//Save calulcating new fish to last

		//Move all fish down one lifecycle
		for j := 1; j < 9; j++ {
			newFishCycles[j-1] = fishCycles[j]
		}

		//Number of dish about to produce new fish
		zeroFish := fishCycles[0]

		//The original zero fish have their lifecycle reset to 6
		newFishCycles[6] = newFishCycles[6] + zeroFish

		//The zerFish all create new fish at lifecycle 8
		newFishCycles[8] = newFishCycles[8] + zeroFish

		fishCycles = newFishCycles
	}

	//All all the fish together
	sum := 0
	for _, fishCount := range fishCycles {
		sum += fishCount
	}

	fmt.Println(fishCycles)
	fmt.Println(sum)
}
