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

type coordinate struct {
	x int
	y int
}

type graphLine struct {
	start  coordinate
	end    coordinate
	points []coordinate
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day5/1/real.txt")
	check(err)

	//PArse the lines
	graphLines := make([]graphLine, 0)
	for _, line := range lines {
		stringCoors := strings.Split(line, " -> ")
		// fmt.Println(stringCoors)

		lineStartEnd := make([]coordinate, 2)
		for i, stringCoor := range stringCoors {
			stringCoordinateSplit := strings.Split(stringCoor, ",")
			var x, errX = strconv.Atoi(stringCoordinateSplit[0])
			check(errX)

			var y, errY = strconv.Atoi(stringCoordinateSplit[1])
			check(errY)

			lineStartEnd[i] = coordinate{x, y}
		}
		graphLine := graphLine{lineStartEnd[0], lineStartEnd[1], make([]coordinate, 0)}
		graphLines = append(graphLines, graphLine)
	}
	// fmt.Println(graphLines)

	//Generate all transit points along the lines
	graphLinesZeroAngle := make([]graphLine, 0)
	//Calculate max x and y while we're at it
	maxX := -1
	maxY := -1
	for _, graphLine := range graphLines {

		startX := graphLine.start.x
		endX := graphLine.end.x
		startY := graphLine.start.y
		endY := graphLine.end.y

		vertical := startX == endX
		horizontal := startY == endY
		if vertical || horizontal {
			points := make([]coordinate, 0)
			if vertical {
				iStart := startY
				iEnd := endY
				if startY > endY {
					iStart = endY
					iEnd = startY
				}

				for i := iStart; i <= iEnd; i++ {
					points = append(points, coordinate{startX, i})
				}

				if iEnd > maxY {
					maxY = iEnd
				}
				if startX > maxX {
					maxX = startX
				}
			} else {
				iStart := startX
				iEnd := endX
				if startX > endX {
					iStart = endX
					iEnd = startX
				}

				for i := iStart; i <= iEnd; i++ {
					points = append(points, coordinate{i, startY})
				}

				if iEnd > maxX {
					maxX = iEnd
				}
				if startY > maxY {
					maxY = startY
				}
			}
			// fmt.Println(points)
			graphLine.points = points
			graphLinesZeroAngle = append(graphLinesZeroAngle, graphLine)
		}
	}
	// fmt.Println(graphLinesZeroAngle)

	// fmt.Println(maxX, maxY)

	//Create map for storing intersection value counts
	intersectionMap := make([][]int, maxX+1)
	for i := range intersectionMap {
		intersectionMap[i] = make([]int, maxY+1)
	}

	//Loop over all generated coordinates and count the intersections
	for _, graphLine := range graphLinesZeroAngle {
		for _, point := range graphLine.points {
			// fmt.Println(point.x, point.y)
			intersectionMap[point.x][point.y] = intersectionMap[point.x][point.y] + 1
		}
	}
	// fmt.Println(intersectionMap)

	//Count number of occurances with a count of two or more
	sumDanger := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			count := intersectionMap[x][y]
			if count >= 2 {
				sumDanger++
			}
		}
	}
	fmt.Println(sumDanger)
}
