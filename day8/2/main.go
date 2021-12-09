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

func calculateSimpleDigit(signal string) int {
	signalLength := len(signal)

	if signalLength == 2 {
		return 1
	} else if signalLength == 4 {
		return 4
	} else if signalLength == 3 {
		return 7
	} else if signalLength == 7 {
		return 8
	}
	return -1
}

func combineSignals(signal1 string, signal2 string) string {
	combinedSignals := signal1
	for _, signal2Rune := range signal2 {
		if !strings.ContainsRune(combinedSignals, signal2Rune) {
			combinedSignals = combinedSignals + string(signal2Rune)
		}
	}
	return combinedSignals
}

func calculateSignalSimilarities(signalToCompareTo string, signalToCompare string) int {
	signalSimilarSum := 0
	for _, signalToCompareRune := range signalToCompare {
		if strings.ContainsRune(signalToCompareTo, signalToCompareRune) {
			signalSimilarSum++
		}
	}

	return signalSimilarSum
}

func areSignalEqual(signal1 string, signal2 string) bool {

	sim1 := calculateSignalSimilarities(signal1, signal2)
	sim2 := calculateSignalSimilarities(signal2, signal1)

	return len(signal1) == len(signal2) && sim1 == sim2 && sim1 == len(signal1) && sim2 == len(signal2)
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day8/2/real.txt")
	check(err)

	outputValues := make([]int, 0)
	for _, line := range lines {
		inputOutput := strings.Split(line, " | ")

		inputs := strings.Split(inputOutput[0], " ")
		outputs := strings.Split(inputOutput[1], " ")

		// fmt.Println(inputs, outputs)

		//Work out all the simple digits
		calculationMap := make(map[int]string)
		for _, input := range inputs {
			simpleDigit := calculateSimpleDigit(input)
			if simpleDigit > 0 {
				calculationMap[simpleDigit] = input
			}
		}
		// fmt.Println(calculationMap)

		//1 is full encapsulated in 3, and three is length 5
		for _, input := range inputs {
			if len(input) == 5 {
				//1 is fully encapsulated so similarities will equal 2
				if calculateSignalSimilarities(input, calculationMap[1]) == 2 {
					calculationMap[3] = input
				}
			}
		}
		// fmt.Println(calculationMap)

		//Calculate 9 by combining 3 with 4
		calculationMap[9] = combineSignals(calculationMap[3], calculationMap[4])
		// fmt.Println(calculationMap)

		//1 is full encapsulated in 3, and three is length 5
		for _, input := range inputs {
			if len(input) == 5 {
				//Check is not 3
				if calculateSignalSimilarities(input, calculationMap[1]) != 2 {
					similar := calculateSignalSimilarities(calculationMap[9], input)
					// fmt.Println(similar)
					if similar == 5 {
						//Must be 5
						calculationMap[5] = input
					} else if similar == 4 {
						//Must be 2
						calculationMap[2] = input
					}
				}
			}
		}
		// fmt.Println(calculationMap)

		for _, input := range inputs {
			if len(input) == 6 {
				//Don't want it to be 9 again
				similar := calculateSignalSimilarities(calculationMap[9], input)
				if similar != 6 {
					similar = calculateSignalSimilarities(calculationMap[1], input)
					if similar == 1 {
						//Must be 6
						calculationMap[6] = input
					} else if similar == 2 {
						//Must be 0
						calculationMap[0] = input
					}
				}
			}
		}
		// fmt.Println(calculationMap)

		//Should have all values, calculate output values
		outputValue := ""
		for _, output := range outputs {
			for key, value := range calculationMap {
				if areSignalEqual(output, value) {
					//Signals match
					outputValue += strconv.Itoa(key)
					break
				}
			}
		}
		value, err := strconv.Atoi(outputValue)
		check(err)
		outputValues = append(outputValues, value)
		// fmt.Println(calculationMap)
		// fmt.Println("Result values: ", outputValues)
	}

	//Sum all the values
	sum := 0
	for _, value := range outputValues {
		sum += value
	}
	fmt.Println(sum)
}
