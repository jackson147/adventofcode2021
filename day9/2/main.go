package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type depthValue struct {
	depthValue int
	lowPoint   bool
	row        int
	col        int
	basin      int
}

func printMap(rows int, cols int, depthMap [][]depthValue) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			basin := depthMap[i][j].basin
			if basin >= 0 {
				fmt.Print(basin, " ")
			} else {
				fmt.Print("x ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day9/2/real.txt")
	check(err)

	//Read in depth map
	rows := len(lines)
	cols := -1
	depthMap := make([][]depthValue, rows)
	for i, line := range lines {
		cols = len(line)
		depthMap[i] = make([]depthValue, cols)
		for j, lineRune := range line {
			lineValue, err := strconv.Atoi(string(lineRune))
			check(err)
			depthMap[i][j] = depthValue{lineValue, false, i, j, -1}
		}
	}
	//

	//Calculate and mark all low-points
	basinCounter := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentDepth := depthMap[i][j]
			currentDepthValue := currentDepth.depthValue

			//Default to true as we assume all are out of bounds
			isUpLower := true
			isDownLower := true
			isLeftLower := true
			isRightLower := true

			//Can check up
			if i > 0 {
				upDepth := depthMap[i-1][j]
				isUpLower = currentDepthValue < upDepth.depthValue
			}

			//Can check down
			if i < rows-1 {
				downDepth := depthMap[i+1][j]
				isDownLower = currentDepthValue < downDepth.depthValue
			}

			if j > 0 {
				leftDepth := depthMap[i][j-1]
				isLeftLower = currentDepthValue < leftDepth.depthValue
			}

			if j < cols-1 {
				rightDepth := depthMap[i][j+1]
				isRightLower = currentDepthValue < rightDepth.depthValue
			}

			currentDepth.lowPoint = isUpLower && isDownLower && isLeftLower && isRightLower
			if currentDepth.lowPoint {
				currentDepth.basin = basinCounter
				basinCounter++
			}
			depthMap[i][j] = currentDepth
		}
	}

	//Get risk levels of low points
	riskLevelSum := 0
	lowPoints := make([]depthValue, 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentDepth := depthMap[i][j]
			if currentDepth.lowPoint {
				riskLevel := currentDepth.depthValue + 1
				riskLevelSum += riskLevel
				// fmt.Println(riskLevel)real

				lowPoints = append(lowPoints, currentDepth)
			}
		}
	}
	// fmt.Println(riskLevelSum)

	//For each low point obtain a list of depths in that basin
	loopCount := 0
	basinNumberChanges := -1
	for basinNumberChanges != 0 {
		basinNumberChanges = 0
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {

				currentDepth := depthMap[i][j]
				currentDepthValue := currentDepth.depthValue

				if i > 0 {
					newDepth := depthMap[i-1][j]
					if newDepth.basin < 0 && newDepth.depthValue != 9 && newDepth.depthValue > currentDepthValue {
						newDepth.basin = currentDepth.basin
						depthMap[i-1][j] = newDepth
						basinNumberChanges++
					}
				}

				if i < rows-1 {
					newDepth := depthMap[i+1][j]
					if newDepth.basin < 0 && newDepth.depthValue != 9 && newDepth.depthValue > currentDepthValue {
						newDepth.basin = currentDepth.basin
						depthMap[i+1][j] = newDepth
						basinNumberChanges++
					}
				}

				if j > 0 {
					newDepth := depthMap[i][j-1]
					if newDepth.basin < 0 && newDepth.depthValue != 9 && newDepth.depthValue > currentDepthValue {
						newDepth.basin = currentDepth.basin
						depthMap[i][j-1] = newDepth
						basinNumberChanges++
					}
				}

				if j < cols-1 {
					newDepth := depthMap[i][j+1]
					if newDepth.basin < 0 && newDepth.depthValue != 9 && newDepth.depthValue > currentDepthValue {
						newDepth.basin = currentDepth.basin
						depthMap[i][j+1] = newDepth
						basinNumberChanges++
					}
				}
			}
		}
		loopCount++
	}
	fmt.Println("COUNT: ", loopCount)

	printMap(rows, cols, depthMap)

	//Calculate sums for each basin
	basinSumMap := make(map[int]int)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentDepth := depthMap[i][j]
			currentBasin := currentDepth.basin

			if currentBasin >= 0 {
				//Increment basin count
				basinSumMap[currentBasin] = basinSumMap[currentBasin] + 1
			}
		}
	}

	finalValues := make([]int, 0)
	for _, value := range basinSumMap {
		finalValues = append(finalValues, value)
	}
	sort.Ints(finalValues)
	//Get final 3 elements
	sum := 0
	for i, value := range finalValues[len(finalValues)-3:] {
		if i == 0 {
			sum = value
		} else {
			sum *= value
		}
	}
	fmt.Println(sum)
}
