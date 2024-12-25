package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Report struct {
	levels []int
}

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	safeReportsCount := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		report := extractReportFromLine(line)
		if isSafe(report) {
			safeReportsCount++
		}
	}
	fmt.Printf("(Part 1) %d reports are safe.\n", safeReportsCount)
}

func part2() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	safeReportsCount := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		report := extractReportFromLine(line)
		if isSafePart2(report) {
			safeReportsCount++
		}
	}
	fmt.Printf("(Part 2) %d reports are safe.\n", safeReportsCount)
}

func extractReportFromLine(line string) Report {
	var currentNumber string
	levels := []int{}
	for _, char := range line {
		if isDigit(char) {
			currentNumber += string(char)
		}
		if char == ' ' || char == '\n' {
			intNum, _ := strconv.Atoi(currentNumber)
			levels = append(levels, intNum)
			currentNumber = ""
		}
	}
	return Report{levels: levels}
}

func isSafe(report Report) bool {
	levels := report.levels
	n := len(levels)
	if n < 2 {
		return true
	}
	isIncreasing := levels[0] < levels[n-1]
	for i := 0; i < n-1; i++ {
		cur := levels[i]
		next := levels[i+1]
		if isIncreasing {
			if next <= cur || next-cur > 3 {
				return false
			}
		} else {
			if next >= cur || cur-next > 3 {
				return false
			}
		}
	}
	return true
}

func isSafePart2(report Report) bool {
	if isSafe(report) {
		return true
	}
	return isSafeIfRemovedOneLevel(report)
}

func isSafeIfRemovedOneLevel(report Report) bool {
	for i := range report.levels {
		newReport := removeLevelFromReport(report, i)
		if isSafe(newReport) {
			return true
		}
	}
	return false
}

func removeLevelFromReport(report Report, i int) Report {
	levels := make([]int, len(report.levels))
	copy(levels, report.levels)
	levels = append(levels[:i], levels[i+1:]...)
	return Report{levels: levels}
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
