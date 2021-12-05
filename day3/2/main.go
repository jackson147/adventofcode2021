package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

type ogBinaryValues struct {
	values []int
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day3/2/real.txt")
	check(err)

	m := make(map[int]binaryLine)
	binM := make(map[int]ogBinaryValues)
	binMLeast := make(map[int]ogBinaryValues)

	//ogBinObjs := make([]ogBinaryValues, 0)
	for i, line := range lines {

		ogValues := make([]int,0)
		sum := 0
		for j := 0; j < len(line); j++ {
			stringValue := string(line[j])
			binarySlice, exits := m[j]
			binaryValue, err := strconv.Atoi(stringValue)
			check(err)

			var values []int
			if !exits {
				values = []int{binaryValue}
			}else{
				values = append(binarySlice.values, binaryValue)
			}
			m[j] = binaryLine{ values, 0, 0}

			sum += binaryValue
			ogValues = append(ogValues, binaryValue)
		}

		//commonBit := -1
		//leastCommonBit := -1
		//if sum < len(ogValues) / 2{
		//	commonBit = 0
		//	leastCommonBit = 1
		//}else{
		//	commonBit = 1
		//	leastCommonBit = 0
		//}

		binM[i] = ogBinaryValues{ogValues}
		binMLeast[i] = ogBinaryValues{ogValues}
	}

	fmt.Println(len(binM))


	//Most common calc
	currentIndex := 0
	for len(binM) > 1 {
		//Calculate sum of index values
		sum := 0
		size := len(binM)
		for _, element := range binM {
			sum += element.values[currentIndex]
		}

		commonBit := 1
		if float32(sum) < float32(size)/float32(2) {
			commonBit = 0
		}

		for key, element := range binM {
			if element.values[currentIndex] != commonBit {
				delete(binM, key)
			}
		}

		time.Sleep(1)
		currentIndex++
	}
	fmt.Println(binM)

	//Most common calc
	currentIndex = 0
	for len(binMLeast) > 1 {
		//Calculate sum of index values
		sum := 0
		size := len(binMLeast)
		for _, element := range binMLeast {
			sum += element.values[currentIndex]
		}

		leastCommonBit := 0
		if float32(sum) < float32(size)/float32(2) {
			leastCommonBit = 1
		}

		for key, element := range binMLeast {
			if element.values[currentIndex] != leastCommonBit {
				delete(binMLeast, key)
			}
		}

		time.Sleep(1)
		currentIndex++
	}
	fmt.Println(binMLeast)

	mostCommonBinaryString := ""
	leastCommonBinaryString := ""

	for  _, element := range binM {
		for _, bit := range element.values {
			mostCommonBinaryString += strconv.Itoa(bit)
		}
	}
	for  _, element := range binMLeast {
		for _, bit := range element.values {
			leastCommonBinaryString += strconv.Itoa(bit)
		}
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