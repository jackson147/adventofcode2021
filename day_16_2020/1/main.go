package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticket struct {
	values []int
}

type bound struct {
	lower int
	upper int
}

type rule struct {
	name string
	bounds []bound
}

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

func getRule(line string) rule {
	splitLine := strings.Split(line, " ")
	ruleName := strings.ReplaceAll(splitLine[0], ":", "")

	limits := splitLine[1:]
	bounds := make([]bound, 0)
	for i, limit := range limits {
		if i % 2 == 0{
			//Must be a bound
			boundStrings := strings.Split(limit, "-")
			lower,_ := strconv.Atoi(boundStrings[0])
			upper,_ := strconv.Atoi(boundStrings[1])

			boundObject := bound{lower, upper}
			bounds = append(bounds, boundObject)
		}else{
			//Must be and/or

			//Ignore for now
		}
	}

	resultRule := rule{ruleName, bounds}

	return resultRule
}

func getTicket(line string) ticket {
	ticketStrings := strings.Split(line, ",")
	var ticketValues []int
	for _, stringValue := range ticketStrings {
		intValue, err := strconv.Atoi(stringValue)
		check(err)
		ticketValues = append(ticketValues, intValue)
	}
	return ticket{ticketValues}
}

func main() {
	lines, err := readLines("/home/rich/git/advent2021/day_16_2020/1/example.txt")
	check(err)

	rules := make([]rule, 0)
	var myTicket ticket
	var otherTickets []ticket

	breakCount := 0
	lineCounter := 0
	for _, line := range lines {

		//Skip this line, it is empty or if it's the first line to read
		if line == "" {
			lineCounter = 0
			breakCount++
			continue
		}

		//Skip the line if it's the start of a new section with a header
		if breakCount > 0 {
			if lineCounter == 0 {
				lineCounter++
				continue
			}
		}

		if breakCount == 0 {
			//Rules
			rules = append(rules, getRule(line))
		}else if breakCount == 1 {
			//My ticket
			myTicket = getTicket(line)
		}else if breakCount == 2 {
			//Other tickets
			otherTickets = append(otherTickets, getTicket(line))
		}
	}

	fmt.Println("RULES: ", rules)
	fmt.Println()
	fmt.Println("My Ticket: ", myTicket)
	fmt.Println()
	fmt.Println("Other tickets: ", otherTickets)
}
