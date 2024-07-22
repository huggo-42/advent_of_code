package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	firstIndex int
	lastIndex  int
	val        int
}

type Symbol struct {
	index      int
	isAsterisk bool
	touchedBy  int
	val1       int
	val2       int
}

func NewNumber(lastIndex, val int) *Number {
	return &Number{
		firstIndex: lastIndex - len(fmt.Sprint(val)),
		lastIndex:  lastIndex - 1,
		val:        val,
	}
}

func NewSymbol(index int, char rune) *Symbol {
	return &Symbol{
		index:      index,
		isAsterisk: char == 42,
		touchedBy:  0,
		val1:       -1,
		val2:       -1,
	}
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var part1sum, part2sum = solve(fileScanner)
	fmt.Println("(Part 1) Sum of all of the part numbers in the engine schematic: ", part1sum)
	fmt.Println("(Part 2) Sum of all of the gear ratios in the engine schematic: ", part2sum)
	readFile.Close()
}

func solve(scanner *bufio.Scanner) (int, int) {
	var part1sum, part2sum, lineIndex = 0, 0, 0
	var numbers [][]*Number
	var symbols [][]*Symbol
	for scanner.Scan() {
		var line = scanner.Text()
		var numbersInLine []*Number
		var symbolsInLine []*Symbol
		var num = ""
		for i, char := range line {
			if isNumber(char) {
				num += string(char)
				if i == 139 {
					var intNum, _ = strconv.Atoi(num)
					numbersInLine = append(numbersInLine, NewNumber(i, intNum))
					num = ""
				}
			} else {
				if isSymbol(char) {
					symbolsInLine = append(symbolsInLine, NewSymbol(i, char))
				}
				if num != "" {
					var intNum, _ = strconv.Atoi(num)
					numbersInLine = append(numbersInLine, NewNumber(i, intNum))
					num = ""
				}
			}
		}
		numbers = append(numbers, numbersInLine)
		symbols = append(symbols, symbolsInLine)
		lineIndex++
	}
	for lineIndex := range symbols {
		var firstLine, lastLine = calcLinesToCheck(lineIndex, len(numbers))
		for i := firstLine; i <= lastLine; i++ {
			for _, number := range numbers[i] {
				for _, symbol := range symbols[lineIndex] {
					if isTouchedBySymbol(number.firstIndex, number.lastIndex, symbol.index) {
						part1sum += number.val
						if symbol.isAsterisk {
							symbol.touchedBy += 1
							if symbol.touchedBy == 1 {
								symbol.val1 = number.val
							} else if symbol.touchedBy == 2 {
								symbol.val2 = number.val
							}
						}
					}
				}
			}
		}
	}
	for _, symbolLine := range symbols {
		for _, symbol := range symbolLine {
			if symbol.isAsterisk && symbol.touchedBy == 2 {
				part2sum += symbol.val1 * symbol.val2
			}
		}
	}
	return part1sum, part2sum
}

func isTouchedBySymbol(firstIndex int, lastIndex int, symbolIndex int) bool {
	return (firstIndex >= symbolIndex-1 && firstIndex <= symbolIndex+1) || (lastIndex <= symbolIndex+1 && lastIndex >= symbolIndex-1)
}

func isNumber(n int32) bool {
	return n > 47 && n < 58
}

func isSymbol(n rune) bool {
	return n != 46
}

func calcLinesToCheck(l int, numsLen int) (int, int) {
	var firstLine, lastLine = l, l
	if l > 0 {
		firstLine = l - 1
	}
	if l < numsLen-1 {
		lastLine = l + 1
	}
	return firstLine, lastLine
}
