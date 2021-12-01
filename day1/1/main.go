package main

import(
	"bufio"
	"fmt"
	"os"
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

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day1/1/real.txt")
	check(err)
	last := 0
	count := 0
	for i, line := range lines {
		fmt.Println(i, line)
		var current, err = strconv.Atoi(line)
		check(err)

		if i > 0 && current > last {
			count++
		}

		last = current
	}
	fmt.Println(count)
}