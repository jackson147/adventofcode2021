package main

import (
	"bufio"
	"fmt"
	"os"
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

func duplicateSmallCaves(route []string) []string {
	duplicateMap := make(map[string]int)
	for _, value := range route {
		if !isUpper(value) && value != "start" && value != "end" {
			duplicateMap[value] = duplicateMap[value] + 1
		}
	}
	duplicates := make([]string, 0)
	for key, value := range duplicateMap {
		if value > 1 {
			duplicates = append(duplicates, key)
		}
	}
	return duplicates
}

type cavePoint struct {
	label       string
	connections []string
	visitCount  int
}

func calculateCaveRoutes(
	caveMap map[string]cavePoint,
	label string,
	currentRoute []string,
	allRoutesRef *([][]string),
) {

	//So we have arrived in the cave, deal with what we do on entry to a cave first
	cave := caveMap[label]

	caveMapCopy := make(map[string]cavePoint)
	for k, v := range caveMap {
		caveMapCopy[k] = v
	}

	currentRouteCopy := make([]string, len(currentRoute))
	for i, v := range currentRoute {
		currentRouteCopy[i] = v
	}

	//We've already tested we're good to be in this cave, so add it
	currentRouteCopy = append(currentRouteCopy, label)
	cave.visitCount = cave.visitCount + 1
	caveMapCopy[label] = cave

	//Is this cave the end? If so we don't bother calculating connections.
	if label == "end" {
		*allRoutesRef = append(*allRoutesRef, currentRouteCopy)

		//We don't consider any new connections so we return
		return
	}

	//For each connection pass back to the same function
	for _, connectionLabel := range cave.connections {
		connection := caveMapCopy[connectionLabel]

		//Is this cave really ok to visit? Must be uppcase or lowecase but visited 0 times
		largeCave := isUpper(connectionLabel)

		//Get duplicates of the current route
		duplicates := duplicateSmallCaves(currentRouteCopy)
		// fmt.Println(duplicates)

		firstSmallCaveVisitOK := !isUpper(connectionLabel) && connection.visitCount < 2 && len(duplicates) == 0
		smallCaveVistedOk := !isUpper(connectionLabel) && connection.visitCount < 1
		smallCaveOk := smallCaveVistedOk || firstSmallCaveVisitOK

		endCave := connectionLabel == "end"
		startCave := connectionLabel == "start"

		//Test cases and compute paths with new map.
		if !startCave && (largeCave || smallCaveOk || endCave) {
			// fmt.Println(connectionLabel, smallCaveNotVisitedOk, smallCaveVistedOk, smallCaveVisited)
			calculateCaveRoutes(
				caveMapCopy,
				connection.label,
				currentRouteCopy,
				allRoutesRef,
			)
		}
	}
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day12/2/real.txt")
	check(err)

	//Read in file
	caveMap := make(map[string]cavePoint)
	for _, line := range lines {
		lineSplit := strings.Split(line, "-")

		//Init cave points if they don't exist
		for _, cavePointLabel := range lineSplit {
			if _, ok := caveMap[cavePointLabel]; !ok {
				caveConnections := make([]string, 0)
				caveMap[cavePointLabel] = cavePoint{cavePointLabel, caveConnections, 0}
			}
		}

		//Add as connections to each other
		caveSource := caveMap[lineSplit[0]]
		caveDest := caveMap[lineSplit[1]]
		caveSource.connections = append(caveSource.connections, caveDest.label)
		caveDest.connections = append(caveDest.connections, caveSource.label)
		caveMap[lineSplit[0]] = caveSource
		caveMap[lineSplit[1]] = caveDest
	}

	// Print inital map
	// for key, value := range caveMap {
	// 	fmt.Println("Cave: ", key, ", Connections: ", value.connections)
	// }

	//Calculate traversal
	allRoutes := make([][]string, 0)
	currentRoute := make([]string, 0)
	calculateCaveRoutes(caveMap, "start", currentRoute, &allRoutes)

	results := make([]string, 0)
	for _, route := range allRoutes {
		result := ""
		for i, value := range route {
			result += value
			if i != len(route)-1 {
				result += ","
			}
		}
		results = append(results, result)
		// fmt.Println(route)
	}
	// fmt.Println(results)

	// resultLines, err := readLines("/home/rich/git/advent2021/day12/2/example_results.txt")
	// resultLineMap := make(map[string]bool)
	// for _, resultLine := range resultLines {
	// 	resultLineMap[resultLine] = false
	// }

	// for _, result := range results {
	// 	delete(resultLineMap, result)
	// }

	// for key := range resultLineMap {
	// 	fmt.Println(key)
	// }

	// for _, route := range allRoutes {
	// 	fmt.Println(route)
	// }

	fmt.Println("TOTAL ROUTES: ", len(allRoutes))
}
