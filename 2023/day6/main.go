package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Race struct {
	time   int
	record int
}

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	times := extractNumbers(line)
	line, _, _ = reader.ReadLine()
	distances := extractNumbers(line)
	solutions := 1
	for i := range times {
		race := Race{times[i], distances[i]}
		solutions = solveRace(race, solutions)
	}
	fmt.Println("(Part 1) I can beat the record in", solutions, "different ways")
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	time := extractOneNumber(line)
	line, _, _ = reader.ReadLine()
	distance := extractOneNumber(line)
	race := Race{time, distance}
	solutions := solveRace(race, 1)
	fmt.Println("(Part 2) I can beat the record in", solutions, "different ways")
}

func solveRace(race Race, solutions int) int {
	firstSolutionFoundAt := 0
	for i := 1; i <= race.time; i++ {
		timeHold, boatSpeed := i, i
		timeMoving := race.time - timeHold
		distance := boatSpeed * timeMoving
		if distance > race.record {
			firstSolutionFoundAt = i
			break
		}
	}
	totalSolutions := (race.time - firstSolutionFoundAt) - firstSolutionFoundAt + 1
	return totalSolutions * solutions
}

func extractNumbers(line []byte) []int {
	numbersArr := []int{}
	currentNumberStr := ""
	for i, char := range line {
		if isNumber(char) {
			currentNumberStr += string(char)
			lastNumber := i == (len(line) - 1)
			if lastNumber {
				numbersArr = append(numbersArr, stringToInt(currentNumberStr))
				break
			}
		} else if len(currentNumberStr) != 0 {
			numbersArr = append(numbersArr, stringToInt(currentNumberStr))
			currentNumberStr = ""
		}
	}
	return numbersArr
}

func extractOneNumber(line []byte) int {
	numberStr := ""
	for _, char := range line {
		if isNumber(char) {
			numberStr += string(char)
		}
	}
	return stringToInt(numberStr)
}

func stringToInt(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}

func isNumber(char byte) bool {
	return char > 47 && char < 58
}
