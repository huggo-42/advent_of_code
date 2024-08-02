package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Seed struct {
	val                int
	solvedForTokenType bool
}

type Range struct {
	start int
	end   int
}

type AlmanacItem struct {
	destinationStart int
	sourceStart      int
	rang             int
}

type TokenType int8

const (
	SeedToSoil            TokenType = 0
	SoilToFertilizer      TokenType = 1
	FertilizerToWater     TokenType = 2
	WaterToLight          TokenType = 3
	LightToTemperature    TokenType = 4
	TemperatureToHumidity TokenType = 5
	HumidityToLocation    TokenType = 6
)

// char codes
const (
	S = 115
	E = 101
	F = 102
	W = 119
	L = 108
	T = 116
	H = 104
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewReader(file)
	line, _, err := scanner.ReadLine()
	seeds := extractSeeds(line)
	currentTokenType := SeedToSoil

	var resetSolvedForTokenType = func() {
		for i := range seeds {
			seeds[i].solvedForTokenType = false
		}
	}

	for err == nil {
		line, _, err = scanner.ReadLine()
		emptyLine := len(line) == 0
		if emptyLine {
			continue
		}
		newTokenType, newTokenFound := checkLineForMarker(line)
		if currentTokenType != newTokenType && newTokenFound {
			currentTokenType = newTokenType
			resetSolvedForTokenType()
			continue
		}
		if !newTokenFound {
			for i := range seeds {
				if seeds[i].solvedForTokenType == true {
					continue
				}
				almanac := extractAlmanacItem(extractNumbers(line))
				almanacStart := almanac.sourceStart
				almanacEnd := almanac.sourceStart + (almanac.rang - 1)
				seedInRange := seeds[i].val >= almanacStart && seeds[i].val <= almanacEnd
				if !seedInRange {
					continue
				}
				seeds[i].val += almanac.destinationStart - almanac.sourceStart
				seeds[i].solvedForTokenType = true
			}
		}
	}
	lowestLocationNumber := findLowestSeed(seeds)
	fmt.Println("(Part 1) The lowest location number is:", lowestLocationNumber)
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewReader(file)
	line, _, err := scanner.ReadLine()
	ranges := extractRanges(line)
	currentTokenType := SeedToSoil
	lowestLocationNumber := ranges[0].start
	solved := []Range{}
	remainers := []Range{}
	var lowestAlmanacStart, highestAlmanacEnd int

	var updateState = func() {
		for _, seedRange := range remainers {
			if outOfBounds(seedRange, lowestAlmanacStart, highestAlmanacEnd) {
				solved = append(solved, seedRange)
			}
		}
		if len(solved) != 0 {
			ranges = solved
		}
		solved = []Range{}
		remainers = []Range{}
	}

	for {
		if err == io.EOF {
			updateState()
			break
		}
		line, _, err = scanner.ReadLine()
		emptyLine := len(line) == 0
		if emptyLine {
			continue
		}
		newTokenType, newTokenFound := checkLineForMarker(line)
		if currentTokenType != newTokenType && newTokenFound {
			currentTokenType = newTokenType
			updateState()
			continue
		}
		if !newTokenFound {
			for _, rang := range ranges {
				almanac := extractAlmanacItem(extractNumbers(line))
				rangeStart := rang.start
				rangeEnd := rang.end
				almanacStart := almanac.sourceStart
				almanacEnd := almanac.sourceStart + (almanac.rang - 1)
				rangeOutOfBounds := rangeStart > almanacEnd || rangeEnd < almanac.sourceStart
				if rangeOutOfBounds {
					continue
				}
				shiftBy := almanac.destinationStart - almanac.sourceStart
				allWithin := rangeStart >= almanacStart && rangeEnd <= almanacEnd
				if allWithin {
					newRange := Range{rangeStart + shiftBy, rangeEnd + shiftBy}
					solved = append(solved, newRange)
				}
				headWithinRange := (rangeStart >= almanacStart && rangeStart <= almanacEnd) && rangeEnd > almanacEnd
				if headWithinRange {
					newRange := Range{rangeStart + shiftBy, almanacEnd + shiftBy}
					solved = append(solved, newRange)
					tail := Range{almanacEnd + 1, rangeEnd}
					remainers = append(remainers, tail)
				}
				tailWithinRange := rangeStart < almanac.sourceStart && (rangeEnd <= almanacEnd && rangeEnd >= almanac.sourceStart)
				if tailWithinRange {
					newRange := Range{almanac.sourceStart + shiftBy, rangeEnd + shiftBy}
					solved = append(solved, newRange)
					head := Range{rangeStart, almanac.sourceStart - 1}
					remainers = append(remainers, head)
				}
				hasHeadAndTail := rangeStart < almanac.sourceStart && rangeEnd > almanacEnd
				if hasHeadAndTail {
					newRange := Range{almanac.sourceStart + shiftBy, almanacEnd + shiftBy}
					solved = append(solved, newRange)
					head := Range{rangeStart, almanac.sourceStart - 1}
					tail := Range{almanacEnd + 1, rangeEnd}
					remainers = append(remainers, head, tail)
				}
				if almanacEnd > highestAlmanacEnd {
					highestAlmanacEnd = almanacEnd
				}
				if almanac.sourceStart < lowestAlmanacStart {
					lowestAlmanacStart = almanac.sourceStart
				}
			}
		}
	}
	lowestLocationNumber = findLowestStartInRanges(&ranges, lowestLocationNumber)
	fmt.Println("(Part 2) The lowest location number is:", lowestLocationNumber)
}

func extractAlmanacItem(line []int) AlmanacItem {
	return AlmanacItem{line[0], line[1], line[2]}
}

func extractRanges(line []byte) []Range {
	numbers := []int{}
	currentNumber := ""
	appendNumber := func() {
		if currentNumber != "" {
			numbers = append(numbers, strToInt(currentNumber))
			currentNumber = ""
		}
	}
	for _, char := range line {
		isNumber := isNumber(char)
		if isNumber {
			currentNumber += string(char)
		} else {
			appendNumber()
		}
	}
	appendNumber() // edge case: last char in line
	ranges := []Range{}
	for i := 0; i < len(numbers); i += 2 {
		ranges = append(ranges, Range{numbers[i], numbers[i] + (numbers[i+1] - 1)})
	}
	return ranges
}

func extractSeeds(line []byte) []Seed {
	numbers := []int{}
	currentNumber := ""
	appendNumber := func() {
		if currentNumber != "" {
			numbers = append(numbers, strToInt(currentNumber))
			currentNumber = ""
		}
	}
	for _, char := range line {
		isNumber := isNumber(char)
		if isNumber {
			currentNumber += string(char)
		} else {
			appendNumber()
		}
	}
	appendNumber() // edge case: last char in line
	seeds := []Seed{}
	for _, num := range numbers {
		seeds = append(seeds, Seed{num, false})
	}
	return seeds
}

func extractNumbers(line []byte) []int {
	numbers := []int{}
	currentNumber := ""
	appendNumber := func() {
		if currentNumber != "" {
			numbers = append(numbers, strToInt(currentNumber))
			currentNumber = ""
		}
	}
	for _, char := range line {
		isNumber := isNumber(char)
		if isNumber {
			currentNumber += string(char)
		} else {
			appendNumber()
		}
	}
	appendNumber() // edge case: last char in line
	return numbers
}

func checkLineForMarker(line []byte) (TokenType, bool) {
	if isNumber(line[0]) {
		return -1, false
	}
	tokenType := SeedToSoil
	switch line[0] {
	case S:
		if line[1] == E {
			tokenType = SeedToSoil
		} else {
			tokenType = SoilToFertilizer
		}
	case F:
		tokenType = FertilizerToWater
	case W:
		tokenType = WaterToLight
	case L:
		tokenType = LightToTemperature
	case T:
		tokenType = TemperatureToHumidity
	case H:
		tokenType = HumidityToLocation
	default:
		tokenType = -1
	}
	return tokenType, tokenType != -1
}

func findLowestSeed(numbers []Seed) int {
	lowest := numbers[0].val
	for _, num := range numbers {
		if num.val < lowest {
			lowest = num.val
		}
	}
	return lowest
}

func findLowestStartInRanges(ranges *[]Range, currentLowest int) int {
	lowestStart := currentLowest
	for _, num := range *ranges {
		if num.start < lowestStart {
			lowestStart = num.start
		}
	}
	return lowestStart
}

func isNumber(char byte) bool {
	return char > 47 && char < 58
}

func strToInt(str string) int {
	intNum, _ := strconv.Atoi(str)
	return intNum
}

func outOfBounds(rang Range, lowest int, highest int) bool {
	if rang.start > highest || rang.end < lowest {
		return true
	}
	return false
}
