package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	distance    int
	previous    string
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
	lines, err := readLines("/home/rich/git/advent2021/day15/1/example.txt")
	// lines, err := readLines("/home/rich/git/advent2021/day15/1/real.txt")
	check(err)

	grid := make([][]riskPoint, len(lines))

	//Load the grid
	for x, line := range lines {
		grid[x] = make([]riskPoint, len(line))
		for y, riskValue := range line {
			stringRiskValue := string(riskValue)
			intRiskValue, err := strconv.Atoi(stringRiskValue)
			check(err)
			key := strconv.Itoa(x) + strconv.Itoa(y)
			grid[x][y] = riskPoint{x, y, key, intRiskValue, make([]*riskPoint, 0), 10000000, ""}
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

	// fmt.Println(grid[5][5].connections)

	//Begin route calculations
	visted := dijkstrasShortestPath(grid, "00")
	printVisted(visted)
	// displayShortestPath("00", visted)

}

func getMinimum(unvisted map[string]*riskPoint) string {
	minKey := "not found"
	minDistance := 10000001 //One higher than default
	for key, risk := range unvisted {
		if risk.distance < minDistance {
			minDistance = risk.distance
			minKey = key
		}
	}
	return minKey
}

func printVisted(visited map[string]*riskPoint) {
	for key, risk := range visited {
		fmt.Println(key, *risk)
	}
}

func dijkstrasShortestPath(grid [][]riskPoint, startNode string) map[string]*riskPoint {

	// # Initialise visited and unvisited lists
	unvisited := make(map[string]*riskPoint) // Declare unvisited list as empty dictionary
	visited := make(map[string]*riskPoint)   // Declare visited list as empty dictionary

	// Add every node to the unvisited list
	for x, row := range grid {
		for y, risk := range row {
			unvisited[risk.key] = &grid[x][y]
		}
	}

	// fmt.Println(unvisited)
	// Set the cost of the start node to 0
	unvisited["00"].distance = 0

	// repeat the following steps until unvisited list is empty
	finished := false
	for !finished {
		if len(unvisited) == 0 {
			finished = true
		} else {
			// Get unvisited node with lowest cost as current node
			currentNodeKey := getMinimum(unvisited)
			currentNode := unvisited[currentNodeKey]
			// Examine neighbours

			for _, neighbour := range currentNode.connections {
				//Only check unvisited neighbours
				if _, ok := visited[neighbour.key]; !ok {
					//Calculate new cost
					cost := unvisited[currentNodeKey].distance + grid[currentNode.x][currentNode.y].distance
					// Check if new cost is less
					if cost < unvisited[neighbour.key].distance {
						unvisited[neighbour.key].distance = cost
						unvisited[neighbour.key].previous = neighbour.key
					}
				}
			}

			// Add current node to visited list
			visited[currentNode.key] = unvisited[currentNode.key]
			// Remove from unvisited list
			delete(unvisited, currentNode.key)
		}
	}

	return visited
}

func displayShortestPath(start string, visited map[string]*riskPoint) {
	for key := range visited {
		fmt.Println(key)
		if key != start {
			current := key
			path := current
			for current != start {
				previous := visited[current].previous
				path = previous + path
				current = visited[current].previous
				fmt.Println(path)
			}
			fmt.Println("Path for: " + key)
			fmt.Println(path)
			fmt.Println("Cost: " + strconv.Itoa(visited[key].distance))
		}
	}
}
