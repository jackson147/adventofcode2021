package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type spoken struct {
	word string
	position int
	count int
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
		valueCounters[line] = spoken{ line, i + 1, 1 }
	}

	numStartingPositions := len(lines)
	lastSpokenWord := lines[numStartingPositions - 1]


	//Start from next turn after reading in list above
	i := numStartingPositions + 1
	for i < 2020 {

		//Map works with string keys
		//key := strconv.Itoa(i)

		//Has the last value been spoken more than once?
		//It will if it exists in the map
		lastSpoken, exists := valueCounters[lastSpokenWord]
		//And if the position is greater than the size of the starting list
		isStartingNumber :=  lastSpoken.position == 1

		if exists && !isStartingNumber {
			//The number is being spoken again. Need to shout the distance from the last position
			check(err)
			justSpokenWord := strconv.Itoa(i - lastSpoken.position)

			//Check if this new number exists already. If so update the position to the current positions
			justSpoken, existsNew := valueCounters[justSpokenWord]
			if existsNew {
				justSpoken.position = i
			}else{//Set a new value
				valueCounters[justSpokenWord] = spoken{justSpokenWord, i, 1}
			}

			lastSpokenWord = justSpokenWord
		}else{
			//The number is new! Shout zero!
			zeroWord := valueCounters["0"]
			zeroWord.position = i
			lastSpokenWord = "0"
		}

		fmt.Println(i, lastSpokenWord)
		i++
	}

	//fmt.Println(valueCounters)
}