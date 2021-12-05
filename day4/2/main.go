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

type boardMarker struct {
	value  int
	marked bool
}

type board struct {
	values [][]boardMarker
}

type coordinate struct {
	boardNumber int
	x           int
	y           int
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day4/2/real.txt")
	check(err)

	valueLookups := make(map[int][]coordinate)
	boards := make([]board, 0)

	var inputs []int

	//Init the board
	currentBoard := board{make([][]boardMarker, 5)}
	for i := range currentBoard.values {
		currentBoard.values[i] = make([]boardMarker, 5)
	}

	boardNumber := 0
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
				currentBoard = board{make([][]boardMarker, 5)}
				for i := range currentBoard.values {
					currentBoard.values[i] = make([]boardMarker, 5)
				}
				boardNumber++
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
				currentBoard.values[rowIndex][newCol] = boardMarker{colValue, false}
				valueLookups[colValue] = append(valueLookups[colValue], coordinate{boardNumber, rowIndex, newCol})
			}

			//Got to increment this for each board entry
			rowIndex++
		}

		if i == len(lines)-1 {
			//Save off the last board
			boards = append(boards, currentBoard)
		}
	}

	// fmt.Println(inputs)
	// fmt.Println(boards)
	// fmt.Println(valueLookups)

	fmt.Println("NUMBER OF BOARDS: ", len(boards))
	// winnerCheckList := make([]bool, len(boards))

	finalWinningBoard := make([]bool, 0)
	finalWinningInput := -1

	winners := make([]bool, len(boards))
	for _, input := range inputs {

		if finalWinningInput > 0 {
			break
		}

		//Mark all board hits
		for _, coordinate := range valueLookups[input] {
			board := (boards[coordinate.boardNumber])
			board.values[coordinate.x][coordinate.y].marked = true
		}

		//Got all the latest winners.
		newWinners := findWinner(boards, winners)

		//Merge with the existing winners
		for i, hasWon := range newWinners {
			if hasWon {
				winners[i] = hasWon
			}
		}

		//Have all bingo cards won?
		countPassed := 0
		for _, hasWon := range winners {
			if hasWon {
				countPassed += 1
			}
		}

		if countPassed == len(boards) {
			finalWinningBoard = newWinners
			finalWinningInput = input

			// fmt.Println(newWinners)
			// fmt.Println(finalWinningInput)
		}

		// fmt.Println(input)
		// fmt.Println(winners)

	}

	finalWinningBoardNumber := -1
	for i, hit := range finalWinningBoard {
		if hit {
			finalWinningBoardNumber = i
			break
		}
	}

	fmt.Println(finalWinningBoardNumber)
	fmt.Println(finalWinningInput)

	winningBoard := boards[finalWinningBoardNumber]
	fmt.Println(finalWinningBoard)
	sumUnmarked := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			marker := winningBoard.values[row][col]
			if !marker.marked {
				sumUnmarked += marker.value
			}
		}
	}
	fmt.Println(sumUnmarked)
	fmt.Println(sumUnmarked * finalWinningInput)

}

func findWinner(boards []board, winners []bool) []bool {
	newWinners := make([]bool, len(winners))
	for boardNum, board := range boards {
		for row := 0; row < 5; row++ {
			colMarkedCount := 0
			for col := 0; col < 5; col++ {
				marker := board.values[row][col]
				if marker.marked {
					colMarkedCount++
				}

				if colMarkedCount == 5 {
					//If not already won
					if !winners[boardNum] {
						newWinners[boardNum] = true
					}
				}
			}
		}

		for col := 0; col < 5; col++ {
			rowMarkedCount := 0
			for row := 0; row < 5; row++ {
				marker := board.values[row][col]
				if marker.marked {
					rowMarkedCount++
				}

				if rowMarkedCount == 5 {
					//If not already won
					if !winners[boardNum] {
						newWinners[boardNum] = true
					}
				}
			}
		}
	}
	return newWinners
}
