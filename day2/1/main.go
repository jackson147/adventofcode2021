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

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day2/1/real.txt")
	check(err)


	position := 0
	depth := 0

	for _, line := range lines {
		cols := strings.Split(line, " ")

		direction := cols[0]
		distance, err := strconv.Atoi(cols[1])
		check(err)
		fmt.Println(direction, distance)

		if direction == "forward" {
			position += distance
		} else if direction == "up" {
			depth -= distance
		} else if direction == "down" {
			depth += distance
		}

		//fmt.Println(i, line)
	}
	fmt.Println(position * depth)
}