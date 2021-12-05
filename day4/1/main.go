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

type board struct {
	values [][]int
}

type coordinate struct {
	x int
	y int
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day4/1/example.txt")
	check(err)

	// valueLookups := make(map[int]coordinate)

	boards := make([]board, 0)

	var inputs []int

	//Init the board
	currentBoard := board{make([][]int, 5)}
	for i := range currentBoard.values {
		currentBoard.values[i] = make([]int, 5)
	}

	rowIndex := 0
	for i, line := range lines {
		// fmt.Println(line)
		if i == 0 {
			//Inputs
			stringInputs := strings.Split(line, ",")
			inputs = make([]int, len(stringInputs))
			for j, stringInput := range stringInputs {
				var intInput, err = strconv.Atoi(stringInput)
				check(err)
				inputs[j] = intInput
			}
		} else if line == "" {
			//Marks the start of a new board, can't be the first line though

			if i > 1 {
				//Save off the old board
				boards = append(boards, currentBoard)
				//Init a new board
				currentBoard = board{make([][]int, 5)}
				for i := range currentBoard.values {
					currentBoard.values[i] = make([]int, 5)
				}
			}

			rowIndex = 0
			continue
		} else {
			boardRowStrings := strings.Split(line, " ")
			colDecrement := 0
			for col, stringValue := range boardRowStrings {
				if stringValue == "" {
					colDecrement++
					continue
				}
				var colValue, err = strconv.Atoi(stringValue)
				check(err)
				newCol := col - colDecrement
				// fmt.Println(rowIndex, newCol, colValue)
				currentBoard.values[rowIndex][newCol] = colValue
			}

			//Got to increment this for each board entry
			rowIndex++
		}

		if i == len(lines)-1 {
			//Save off the last board
			boards = append(boards, currentBoard)
		}
	}

	fmt.Println(inputs)
	fmt.Println(boards)
}
