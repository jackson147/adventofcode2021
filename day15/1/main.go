package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/RyanCarrier/dijkstra"
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

func copyMap(toCopy map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range toCopy {
		newMap[key] = value
	}
	return newMap
}

func combineMaps(toAddTo map[string]int, toAdd map[string]int) {
	for key, value := range toAdd {
		toAddTo[key] = toAddTo[key] + value
	}
}

type riskPoint struct {
	x           int
	y           int
	key         string
	risk        int
	connections []*riskPoint
	id          int
}

func printRiskGrid(grid [][]riskPoint) {
	for _, row := range grid {
		for _, point := range row {
			fmt.Print(point.risk)
		}
		fmt.Println()
	}
}

func main() {
	// lines, err := readLines("/home/rich/git/advent2021/day15/1/example.txt")
	lines, err := readLines("/home/rich/git/advent2021/day15/1/real.txt")
	check(err)

	grid := make([][]riskPoint, len(lines))

	//Load the grid
	id := 0
	for x, line := range lines {
		grid[x] = make([]riskPoint, len(line))
		for y, riskValue := range line {
			stringRiskValue := string(riskValue)
			intRiskValue, err := strconv.Atoi(stringRiskValue)
			check(err)
			key := strconv.Itoa(x) + strconv.Itoa(y)
			grid[x][y] = riskPoint{x, y, key, intRiskValue, make([]*riskPoint, 0), id}
			id++
		}
	}

	// printRiskGrid(grid)

	//Make connections to other grid references
	for x, row := range grid {
		for y, _ := range row {
			//We can go negative x
			if x > 0 {
				grid[x][y].connections = append(grid[x][y].connections, &grid[x-1][y])
			}

			//We can go postive x
			if x < len(row)-1 {
				grid[x][y].connections = append(grid[x][y].connections, &grid[x+1][y])
			}

			//We can go negative y
			if y > 0 {
				grid[x][y].connections = append(grid[x][y].connections, &grid[x][y-1])
			}

			//We can go postive y
			if y < len(row)-1 {
				grid[x][y].connections = append(grid[x][y].connections, &grid[x][y+1])
			}
		}
	}

	printRiskGrid(grid)

	graph := dijkstra.NewGraph()

	for _, row := range grid {
		for _, risk := range row {
			// fmt.Println("Adding Vertex ", risk.id)
			graph.AddVertex(risk.id)
		}
	}

	for _, row := range grid {
		for _, risk := range row {
			for _, connection := range risk.connections {
				// fmt.Println(risk.id, connection.id, int64(connection.risk))
				graph.AddArc(risk.id, connection.id, int64(connection.risk))
			}
		}
	}

	best, err := graph.Shortest(0, len(lines)*len(lines)-1)
	check(err)
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
}
