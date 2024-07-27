package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	COLON        = 58
	VERTICAL_BAR = 124
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewReader(file)
	points := 0
	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		points += calcPointsInScratchcard(line)
	}
	fmt.Println("(Part 1) All those scratchcards are worth", points, "points")
}

func part2() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewReader(file)
	scratchcardsMap := map[int]int{}
	scratchcards, scratchcardIndex := 0, 1
	for {
		scratchcardsMap[scratchcardIndex] += 1
		line, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		correctNumbers := countCorrectNumbers(line)
		makeScratchcardCopies(scratchcardIndex, correctNumbers, scratchcardsMap)
		scratchcards += scratchcardsMap[scratchcardIndex]
		scratchcardIndex++
	}
	fmt.Println("(Part 2) The Elf will end up with", scratchcards, "scratchcards")
}

func countCorrectNumbers(line []byte) int {
	winningNumbers := []int{}
	myNumbers := []int{}
	pastGameIdMarker, pastWinningNumbersMarker := false, false
	numString := ""
	for i, char := range line {
		if char == COLON {
			pastGameIdMarker = true
			continue
		} else if char == VERTICAL_BAR {
			pastWinningNumbersMarker = true
			continue
		}
		if !pastGameIdMarker {
			continue
		}
		isNumber, isLastIndex := isNumber(char), isLastIndex(i, len(line))
		if isNumber {
			numString += string(char)
		}
		// edge case
		if isNumber && isLastIndex {
			intNum, _ := strconv.Atoi(numString)
			myNumbers = append(myNumbers, intNum)
			numString = ""
			continue
		}
		if !isNumber && !isStringEmpty(numString) {
			intNum, _ := strconv.Atoi(numString)
			if !pastWinningNumbersMarker {
				winningNumbers = append(winningNumbers, intNum)
			}
			if pastWinningNumbersMarker {
				myNumbers = append(myNumbers, intNum)
			}
			numString = ""
		}
	}
	correctNumbers := 0
	for _, myNum := range myNumbers {
		for _, winningNum := range winningNumbers {
			if myNum == winningNum {
				correctNumbers++
			}
		}
	}
	return correctNumbers
}

func calcPointsInScratchcard(line []byte) int {
	var points int
	correctNumbers := countCorrectNumbers(line)
	if correctNumbers == 0 {
		points = 0
	} else if correctNumbers == 1 {
		points = 1
	} else if correctNumbers > 1 {
		points = pow2(correctNumbers - 1)
	}
	return points
}

func isNumber(char byte) bool {
	return char > 47 && char < 58
}

func pow2(num int) int {
	factor := 2
	for i := 1; i < num; i++ {
		factor *= 2
	}
	return factor
}

func isLastIndex(index int, length int) bool {
	return index == (length - 1)
}

func isStringEmpty(str string) bool {
	return len(str) == 0
}

func makeScratchcardCopies(scratchcardIndex int, copiesNum int, scratchcardsMap map[int]int) {
	for i := 1; i <= copiesNum; i++ {
		scratchcardsMap[scratchcardIndex+i] += 1 * scratchcardsMap[scratchcardIndex]
	}
}
