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

func printCache(cache map[string]map[string]int) {
	for key, value := range cache {
		fmt.Println(key, value)
	}
}

var polymerMap map[string]string

func main() {
	// lines, err := readLines("/home/rich/git/advent2021/day14/2/example.txt")
	lines, err := readLines("/home/rich/git/advent2021/day14/2/real.txt")
	check(err)

	instructionLine := ""

	polymerMap = make(map[string]string)

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			instructionLine = line
		} else {
			polymers := strings.Split(line, " -> ")
			polymerMap[polymers[0]] = polymers[1]
		}
	}

	// instructionLine = "CH"
	fmt.Println(instructionLine)
	fmt.Println(polymerMap)

	pairs := make([][]string, 0)
	currentPair := make([]string, 0)
	// resultMap := make(map[string]int)

	costMap := make(map[string]int)
	for _, instructionRune := range instructionLine {
		currentPair = append(currentPair, string(instructionRune))
		if len(currentPair) == 2 {
			pairs = append(pairs, currentPair)
			currentPair = currentPair[1:]
		}
		costMap[string(instructionRune)] = costMap[string(instructionRune)] + 1
	}

	totalIterations := 40

	//Prime the cache
	cache = make(map[string]map[string]int)
	for i := 0; i < totalIterations; i++ {
		for polyPair, _ := range polymerMap {
			cacheCost := make(map[string]int)
			polyOne := string(polyPair[0])
			polyTwo := string(polyPair[1])
			calculateCost(cacheCost, polyOne, polyTwo, i)

			cacheKey := polyOne + polyTwo + strconv.Itoa(i)
			cache[cacheKey] = copyMap(cacheCost)
		}
	}

	for _, pair := range pairs {
		// fmt.Println(i)
		calculateCost(costMap, pair[0], pair[1], totalIterations-1)
	}

	fmt.Println(costMap)
	values := make([]int, 0)
	for _, value := range costMap {
		values = append(values, value)
	}
	printCache(cache)
	min, max := minMax(values)
	fmt.Println(max - min)
}

var cache map[string]map[string]int

func calculateCost(costMap map[string]int, polyOne string, polyTwo string, value int) map[string]int {
	//Cache map to save
	cacheCost := make(map[string]int)

	// fmt.Println(polyOne, polyTwo)
	polyPair := polyOne + polyTwo

	//Get the new polymer from the inputs
	newPolymer := polymerMap[polyPair]
	//Always add the new polymer to the costmap
	costMap[newPolymer] = costMap[newPolymer] + 1
	cacheCost[newPolymer] = cacheCost[newPolymer] + 1

	//We're all done, return the cost map
	if value == 0 {
		// fmt.Println("RETURN")
		return costMap
	}

	//Calculate the cost of the new polymers
	cacheKey := polyOne + newPolymer + strconv.Itoa(value-1)
	if cachedResult, ok := cache[cacheKey]; ok {
		combineMaps(costMap, cachedResult)
	} else {
		calculateCost(costMap, polyOne, newPolymer, value-1)
	}
	// calculateCost(costMap, polyOne, newPolymer, value-1)

	cacheKey = newPolymer + polyTwo + strconv.Itoa(value-1)
	if cachedResult, ok := cache[cacheKey]; ok {
		combineMaps(costMap, cachedResult)
	} else {
		calculateCost(costMap, newPolymer, polyTwo, value-1)
	}
	// calculateCost(costMap, newPolymer, polyTwo, value-1)

	//We'll only get to this point when the the function calls above start reaching value == 0
	return costMap
}
