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
	// lines, err := readLines("/home/rich/git/advent2021/day15/2/example.txt")
	lines, err := readLines("/home/rich/git/advent2021/day15/2/real.txt")
	check(err)

	gridSize := len(lines)
	fullGridSize := 5 * gridSize

	//Init the grid
	grid := make([][]riskPoint, fullGridSize)
	for x := 0; x < fullGridSize; x++ {
		grid[x] = make([]riskPoint, fullGridSize)
	}

	//Load the grid
	id := 0
	for x, line := range lines {
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

	//Repeats 5 times in X direction
	for i := 0; i < 5; i++ {
		//Repeats 5 times in Y direction
		for j := 0; j < 5; j++ {
			//Loops over the previous grid in the X direction
			for x := 0; x < gridSize; x++ {
				//Loops over the previous grid in the Y direction
				for y := 0; y < gridSize; y++ {

					//New grid coordinates
					newX := x + i*gridSize
					newY := y + j*gridSize
					// fmt.Println(newX, x, i*5)
					// fmt.Println(newY, y, j*5)

					// fmt.Println(x, y, newX, newY)

					//If i ==0 && j==0 don't read from a source
					sourceX := -1
					sourceY := -1
					if i == 0 && j == 0 {
						sourceX = x
						sourceY = y
						//If y==0 read from x-10, but y=0
					} else if j == 0 {
						sourceX = newX - gridSize
						sourceY = newY
						//If x==0 read from y-10, but x=0
					} else if i == 0 {
						sourceX = newX
						sourceY = newY - gridSize
						//If y>0 read from y-10
					} else {
						sourceX = newX - gridSize
						sourceY = newY
					}

					// fmt.Println(newX, newY, sourceX, sourceY)

					newRisk := grid[sourceX][sourceY].risk
					if !(i == 0 && j == 0) {
						newRisk = grid[sourceX][sourceY].risk + 1
						if newRisk > 9 {
							newRisk = 1
						}
					}

					key := strconv.Itoa(newX) + strconv.Itoa(newY)

					grid[newX][newY] = riskPoint{newX, newY, key, newRisk, make([]*riskPoint, 0), id}
					id++
				}
			}
		}
	}
	fmt.Println(id)
	printRiskGrid(grid)

	//Make connections to other grid references
	for x, row := range grid {
		for y := range row {
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

	// printRiskGrid(grid)

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

	lenGrid := len(grid) - 1
	lowestId := grid[0][0].id
	highestId := grid[lenGrid][lenGrid].id
	best, err := graph.Shortest(lowestId, highestId)
	check(err)
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)
}
