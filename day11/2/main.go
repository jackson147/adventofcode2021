package main

import (
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

type octopus struct {
	energy  int
	flashed bool
}

func printOctoGrid(grid [][]octopus) {
	for _, row := range grid {
		for _, octo := range row {
			fmt.Print(octo.energy, " ")
		}
		fmt.Println("")
	}
}

func increaseAllOctoEnergy(grid [][]octopus) [][]octopus {
	for i, row := range grid {
		for j, _ := range row {
			grid = increaseOctoEnergy(grid, i, j)
		}
	}
	return grid
}

func resetAllFlashed(grid [][]octopus) ([][]octopus, int) {
	flashedSum := 0
	for i, row := range grid {
		for j, octo := range row {
			if octo.flashed {
				octo.energy = 0
				octo.flashed = false
				grid[i][j] = octo
				flashedSum++
			}
		}
	}
	return grid, flashedSum
}

func increaseOctoEnergy(grid [][]octopus, i int, j int) [][]octopus {
	// fmt.Println(i, j)
	octo := grid[i][j]
	octo.energy = octo.energy + 1
	grid[i][j] = octo
	return grid
}

func flashOctopus(grid [][]octopus, i int, j int) [][]octopus {

	//10 or more to flash
	if grid[i][j].flashed != true && grid[i][j].energy > 9 {

		//This octopus has flashed!
		grid[i][j].flashed = true

		//Calculate where we can look in the 2d array
		northOk := i > 0
		eastOk := j < len(grid[0])-1
		southOk := i < len(grid)-1
		westOk := j > 0

		northInc := i - 1
		eastInc := j + 1
		southInc := i + 1
		westInc := j - 1

		//Increment all energy in allowed directions
		//Flash if appropriate!

		//N
		if northOk {
			grid = increaseOctoEnergy(grid, northInc, j)
			grid = flashOctopus(grid, northInc, j)
		}
		//NE
		if northOk && eastOk {
			grid = increaseOctoEnergy(grid, northInc, eastInc)
			grid = flashOctopus(grid, northInc, eastInc)
		}
		//E
		if eastOk {
			grid = increaseOctoEnergy(grid, i, eastInc)
			grid = flashOctopus(grid, i, eastInc)
		}
		//SE
		if southOk && eastOk {
			grid = increaseOctoEnergy(grid, southInc, eastInc)
			grid = flashOctopus(grid, southInc, eastInc)
		}
		//S
		if southOk {
			grid = increaseOctoEnergy(grid, southInc, j)
			grid = flashOctopus(grid, southInc, j)
		}
		//SW
		if southOk && westOk {
			// fmt.Println(i, j, southInc, westInc)
			grid = increaseOctoEnergy(grid, southInc, westInc)
			grid = flashOctopus(grid, southInc, westInc)
		}
		//W
		if westOk {
			grid = increaseOctoEnergy(grid, i, westInc)
			grid = flashOctopus(grid, i, westInc)
		}
		//NW
		if northOk && westOk {
			grid = increaseOctoEnergy(grid, northInc, westInc)
			grid = flashOctopus(grid, northInc, westInc)
		}

	}
	return grid
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day11/2/real.txt")
	check(err)

	//Read in file
	octoGrid := make([][]octopus, len(lines))
	for i, line := range lines {
		octoGrid[i] = make([]octopus, len(line))
		for j, octoValueRune := range line {
			octoValueString := string(octoValueRune)
			octoValue, err := strconv.Atoi(octoValueString)
			check(err)
			octoGrid[i][j] = octopus{octoValue, false}
		}
	}
	// fmt.Println(octoGrid)
	// printOctoGrid(octoGrid)

	counter := 0
	numberFlashed := 0
	//Total iterations
	for numberFlashed != 100 {
		//1 increase all octo values by 1
		octoGrid = increaseAllOctoEnergy(octoGrid)

		//Loop through and check for need to flash octopuses (lol)
		for i, row := range octoGrid {
			for j, _ := range row {
				octoGrid = flashOctopus(octoGrid, i, j)
			}
		}
		octoGrid, numberFlashed = resetAllFlashed(octoGrid)

		printOctoGrid(octoGrid)
		fmt.Println("")
		counter++
	}
	fmt.Println("STEP ALL FLASH: ", counter)
}
