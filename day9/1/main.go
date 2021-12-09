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

type depthValue struct {
	depthValue int
	lowPoint   bool
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day9/1/real.txt")
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
			depthMap[i][j] = depthValue{lineValue, false}
		}
	}
	//

	//Calculate and mark all low-points
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
			depthMap[i][j] = currentDepth
		}
	}

	//Get risk levels of low points
	riskLevelSum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentDepth := depthMap[i][j]
			if currentDepth.lowPoint {
				riskLevel := currentDepth.depthValue + 1
				riskLevelSum += riskLevel
				// fmt.Println(riskLevel)real
			}
		}
	}
	fmt.Println(riskLevelSum)
}
