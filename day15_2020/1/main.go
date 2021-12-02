package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type spoken struct {
	word             string
	position         int
	previousPosition int
	count            int
}

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

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day15_2020/1/example.txt")
	check(err)

	valueCounters := make(map[string]spoken)

	for i, line := range lines {
		check(err)
		valueCounters[line] = spoken{line, i + 1, -1, 1}
	}

	numStartingPositions := len(lines)
	lastSpokenWord := lines[numStartingPositions-1]

	//Start from next turn after reading in list above
	i := numStartingPositions + 1
	for i <= 30000000 {

		//Has the last value been spoken more than once?
		//It will if it exists in the map
		lastSpoken := valueCounters[lastSpokenWord]

		// fmt.Println("Last number spoken", lastSpokenWord)
		// fmt.Println(valueCounters)

		if lastSpoken.count == 1 {
			//The number was spoken for the first time! Speak ZERO for the new word
			zeroWord := valueCounters["0"]
			zeroWord.count = zeroWord.count + 1
			zeroWord.previousPosition = zeroWord.position
			zeroWord.position = i

			valueCounters["0"] = zeroWord
			lastSpokenWord = "0"
		} else if lastSpoken.count > 1 {
			//The number is being spoken again. Need to shout the distance from the last position
			diff := lastSpoken.position - lastSpoken.previousPosition
			justSpokenWord := strconv.Itoa(diff)
			// fmt.Println("DIFFERENCE:", diff, i, lastSpoken.position)

			justSpoken := valueCounters[justSpokenWord]
			justSpoken.count = justSpoken.count + 1
			justSpoken.previousPosition = justSpoken.position
			justSpoken.position = i
			valueCounters[justSpokenWord] = justSpoken

			lastSpokenWord = justSpokenWord
		}

		// fmt.Println(i, lastSpokenWord)
		i++
	}
	fmt.Println(lastSpokenWord)
}
