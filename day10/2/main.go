package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func printRuneSlice(runeSlice []rune) {
	for _, runeElement := range runeSlice {
		fmt.Print(string(runeElement))
	}
	fmt.Println("")
}

var syntaxOppositesMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

var syntaxErrorPoints = map[rune]int{
	')': 3,
	'}': 1197,
	']': 57,
	'>': 25137,
}

var autoCompletePoints = map[rune]int{
	')': 1,
	'}': 3,
	']': 2,
	'>': 4,
}

func isCorrupt(line string) (bool, rune) {
	openingRunes := make([]rune, 0)
	for _, syntaxRune := range line {
		if syntaxRune == '(' || syntaxRune == '[' || syntaxRune == '<' || syntaxRune == '{' {
			//Opening rune
			openingRunes = append(openingRunes, syntaxRune)
		} else {
			//Closing rune

			// printRuneSlice(openingRunes)
			// fmt.Println(string(syntaxRune))
			// fmt.Println()

			//To add a closing rune, it must match the last opening rune
			lastOpeningRune := openingRunes[len(openingRunes)-1]
			if syntaxRune != syntaxOppositesMap[lastOpeningRune] {
				//Syntax error
				return true, syntaxRune
			}

			//Remove the last element from the opening runes, we have closed it
			openingRunes = openingRunes[:len(openingRunes)-1]
		}
	}

	//Return space as an empty rune
	return false, ' '
}

func getClosingSequence(incompleteLine string) []rune {
	//This function depends on the incoming line having no syntax errors
	openingRunes := make([]rune, 0)
	for _, syntaxRune := range incompleteLine {
		if syntaxRune == '(' || syntaxRune == '[' || syntaxRune == '<' || syntaxRune == '{' {
			//Opening rune
			openingRunes = append(openingRunes, syntaxRune)
		} else {
			//Closing rune
			//Remove the last element from the opening runes, we have closed it
			openingRunes = openingRunes[:len(openingRunes)-1]
		}
	}

	// fmt.Println("REMAINING OPENING RUNES: ")
	// printRuneSlice(openingRunes)

	//At this point we should only have the remaining incomplete runes to close
	//Iterate down to get each closing rune
	closingRunes := make([]rune, len(openingRunes))
	counter := 0
	for i := len(openingRunes) - 1; i >= 0; i-- {
		closingRune := syntaxOppositesMap[openingRunes[i]]
		closingRunes[counter] = closingRune
		counter++
	}

	//Return space as an empty rune
	return closingRunes
}

func calculateAutoCompleteScore(symbols []rune) int {
	score := 0
	for _, symbol := range symbols {
		score *= 5
		score += autoCompletePoints[symbol]
	}
	return score
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day10/2/real.txt")
	check(err)

	incompleteLines := make([]string, 0)
	corruptRunes := make([]string, 0)
	syntaxErrorPointSum := 0
	for _, line := range lines {
		isCorrupt, corruptRune := isCorrupt(line)
		if isCorrupt {
			corruptRunes = append(corruptRunes, string(corruptRune))
			syntaxErrorPointSum += syntaxErrorPoints[corruptRune]
		} else {
			incompleteLines = append(incompleteLines, line)
		}
	}
	fmt.Println("# Corrupt lines: ", len(corruptRunes))
	fmt.Println("Corrupt runes: ", corruptRunes)
	fmt.Println("Corrupt points score: ", syntaxErrorPointSum)

	scores := make([]int, 0)
	for _, line := range incompleteLines {
		closingRunes := getClosingSequence(line)
		// printRuneSlice(closingRunes)
		score := calculateAutoCompleteScore(closingRunes)
		scores = append(scores, score)
	}
	sort.Ints(scores)
	middleValueIndex := (len(scores) / 2) //Base index 0 means we don't need to add one
	fmt.Println("Auto complete winner: ", scores[middleValueIndex])
}
