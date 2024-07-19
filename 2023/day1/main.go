package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part1sum int
	var part2sum int
	for fileScanner.Scan() {
		part1sum += calcCalibrationValueForlinePart1(fileScanner.Text())
		part2sum += calcCalibrationValueForlinePart2(fileScanner.Text())
	}
	readFile.Close()
	fmt.Println("Sum of all of the calibration values: ", part1sum)
	fmt.Println("Correct sum of all of the calibration values: ", part2sum)
}

func calcCalibrationValueForlinePart1(line string) int {
	first, last := -1, -1
	firstPos, lastPos := len(line), -1
	for idx, char := range line {
		if isNumber(char) {
			intChar, _ := strconv.Atoi(string(char))
			if idx < firstPos {
				firstPos = idx
				first = intChar
			}
			if idx > lastPos {
				lastPos = idx
				last = intChar
			}
		}
	}
	calibrationValue, _ := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(last))
	return calibrationValue
}

func calcCalibrationValueForlinePart2(line string) int {
	first, last := -1, -1
	firstPos, lastPos := len(line), -1
	for idx, char := range line {
		if isNumber(char) {
			intChar, _ := strconv.Atoi(string(char))
			if idx < firstPos {
				firstPos = idx
				first = intChar
			}
			if idx > lastPos {
				lastPos = idx
				last = intChar
			}
		}
	}
	for _, numStr := range numStrings {
		if strings.Contains(line, numStr) {
			strIdx := strings.Index(line, numStr)
			if strIdx < firstPos {
				firstPos = strIdx
				first = stringToInt(numStr)
			}
			lastStrIdx := strings.LastIndex(line, numStr)
			if lastStrIdx > lastPos {
				lastPos = lastStrIdx
				last = stringToInt(numStr)
			}
		}
	}
	calibrationValue, _ := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(last))
	return calibrationValue
}

func isNumber(n int32) bool {
	return n > 47 && n < 58
}

var numStrings = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func stringToInt(numStr string) int {
	switch numStr {
	case numStrings[0]:
		return 0
	case numStrings[1]:
		return 1
	case numStrings[2]:
		return 2
	case numStrings[3]:
		return 3
	case numStrings[4]:
		return 4
	case numStrings[5]:
		return 5
	case numStrings[6]:
		return 6
	case numStrings[7]:
		return 7
	case numStrings[8]:
		return 8
	case numStrings[9]:
		return 9
	}
	return -1
}
