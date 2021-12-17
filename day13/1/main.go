package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type foldInstruction struct {
	axis  string
	value int
}

func printGrid(grid [][]bool) {
	maxX := len(grid)
	maxY := len(grid[0])
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func calculateVisibleDots(grid [][]bool) int {
	maxX := len(grid)
	maxY := len(grid[0])

	sumDots := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[x][y] {
				sumDots++
			}
		}
	}
	return sumDots
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day13/1/real.txt")
	check(err)

	instructionBreak := false
	minX := 999999
	maxX := -1
	minY := 999999
	maxY := -1

	coordinates := make([][]int, 0)
	foldInstructions := make([]foldInstruction, 0)
	for _, line := range lines {
		if line == "" {
			instructionBreak = true
			//Skip this line
			continue
		}

		if !instructionBreak {
			lineSplit := strings.Split(line, ",")
			x, err := strconv.Atoi(lineSplit[0])
			check(err)
			y, err := strconv.Atoi(lineSplit[1])
			check(err)

			if x < minX {
				minX = x
			}
			if y < minY {
				minY = y
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}

			coordinate := []int{x, y}
			coordinates = append(coordinates, coordinate)
		} else {
			//Process fold instructions
			lineSplit := strings.Split(line, "fold along ")
			foldLineSplit := strings.Split(lineSplit[1], "=")
			axis := foldLineSplit[0]
			value, err := strconv.Atoi(foldLineSplit[1])
			check(err)

			foldInstructions = append(foldInstructions, foldInstruction{axis, value})
		}
	}

	// fmt.Println(coordinates)

	//Init the grid
	grid := make([][]bool, maxX+1)
	for i := range grid {
		grid[i] = make([]bool, maxY+1)
	}

	// fmt.Println(maxX, maxY)
	for _, coordinate := range coordinates {
		// fmt.Println(coordinate)
		grid[coordinate[0]][coordinate[1]] = true
	}

	// printGrid(grid)

	// fmt.Println(foldInstructions)

	// fmt.Println(foldInstructions[0])
	newGrid := processFoldInstruction(grid, foldInstructions[0])
	// fmt.Println()
	// printGrid(newGrid)

	fmt.Println("Number dots: ", calculateVisibleDots(newGrid))

	// newGrid = processFoldInstruction(newGrid, foldInstructions[1])
	// fmt.Println()
	// printGrid(newGrid)

}

func processFoldInstruction(grid [][]bool, instruction foldInstruction) [][]bool {
	foldLine := instruction.value
	var newGrid [][]bool
	maxX := len(grid)
	maxY := len(grid[0])
	if instruction.axis == "y" {
		//Removes the fold line as an option
		newGridMaxY := maxY - foldLine - 1

		// fmt.Println(maxX, maxY, newGridMaxY)

		//Retains the X size
		//Init the grid
		newGrid = make([][]bool, maxX)
		for row := range newGrid {
			newGrid[row] = make([]bool, newGridMaxY)
		}

		//Tranpose all old points
		for row := 0; row < maxX; row++ {
			for col := 0; col < newGridMaxY; col++ {
				newGrid[row][col] = grid[row][col]
			}
		}
		// printGrid(newGrid)

		for row := 0; row < maxX; row++ {
			yCounter := 0
			for col := newGridMaxY + 1; col < maxY; col++ {
				mappedY := newGridMaxY - 1 - yCounter
				marked := grid[row][col]
				if marked {
					newGrid[row][mappedY] = marked
				}
				yCounter++
			}
		}
	} else {
		//Must be x
		//Removes the fold line as an option
		newGridMaxX := maxX - foldLine

		fmt.Println(maxX, maxY, newGridMaxX)

		//Retains the X size
		//Init the grid
		newGrid = make([][]bool, newGridMaxX)
		for row := range newGrid {
			newGrid[row] = make([]bool, maxY+1)
		}

		// Tranpose all old points
		for row := 0; row < newGridMaxX; row++ {
			for col := 0; col < maxY; col++ {
				newGrid[row][col] = grid[row][col]
			}
		}
		// printGrid(newGrid)

		xCounter := 0
		lastX := -1
		for row := newGridMaxX + 1; row < maxX; row++ {
			for col := 0; col < maxY; col++ {
				mappedX := newGridMaxX - 1 - xCounter

				if row != lastX {
					fmt.Println(row, col, mappedX, foldLine, newGridMaxX)
					lastX = row
				}

				marked := grid[row][col]
				if marked {
					newGrid[mappedX][col] = marked
				}
			}
			xCounter++
		}
	}
	return newGrid
}
