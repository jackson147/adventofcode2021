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

type binaryLine struct {
	values []int
	commonBit int
	leastCommon int
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day3/1/real.txt")
	check(err)

	m := make(map[int]binaryLine)

	for _, line := range lines {
		//fmt.Println(line)
		for i := 0; i < len(line); i++ {
			stringValue := string(line[i])
			binarySlice, exits := m[i]
			binaryValue, err := strconv.Atoi(stringValue)
			check(err)

			var values []int
			if !exits {
				values = []int{binaryValue}
			}else{
				values = append(binarySlice.values, binaryValue)
			}
			m[i] = binaryLine{values, 0, 0}
		}
	}


	for key, element := range m {
		sum := 0
		for _,binaryValue := range element.values {
			//fmt.Println(binaryValue)
			sum += binaryValue
		}
		//fmt.Println(key, sum)

		if sum < len(element.values)/2 {
			element.commonBit = 0
			element.leastCommon = 1
		}else{
			element.commonBit = 1
			element.leastCommon = 0
		}

		m[key] = element
	}

	fmt.Println(m)

	//Sort the keys in the map
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	mostCommonBinaryString := ""
	leastCommonBinaryString := ""
	for _, key := range keys {
		element := m[key]

		mostCommonStringBit := strconv.Itoa(element.commonBit)
		leastCommonStringBit := strconv.Itoa(element.leastCommon)

		mostCommonBinaryString += mostCommonStringBit
		leastCommonBinaryString += leastCommonStringBit
	}

	parsedMostCommon, err := strconv.ParseInt(mostCommonBinaryString, 2, 64)
	parsedLeastCommon, err := strconv.ParseInt(leastCommonBinaryString, 2, 64)
	check(err)

	fmt.Println(mostCommonBinaryString)
	fmt.Println(leastCommonBinaryString)

	fmt.Println(parsedMostCommon)
	fmt.Println(parsedLeastCommon)

	fmt.Println(parsedMostCommon * parsedLeastCommon)
}