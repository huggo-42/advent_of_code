package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	leftArr := []int{}
	rightArr := []int{}
	for {
		line, err := reader.ReadString('\n')
		leftNum, rightNum := extractNumbersFromLine(line)
		if err != nil {
			break
		}
		leftArr = insertMaintainingOrder(leftArr, leftNum)
		rightArr = insertMaintainingOrder(rightArr, rightNum)
	}
	totalDistance := calcTotalDistance(leftArr, rightArr)
	fmt.Printf("(Part 1) The total distance between the left list and the right list is %d\n", totalDistance)
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	leftArr := []int{}
	rightMap := make(map[int]int)
	for {
		line, err := reader.ReadString('\n')
		leftNum, rightNum := extractNumbersFromLine(line)
		if err != nil {
			break
		}
		leftArr = append(leftArr, leftNum)
		rightMap[rightNum] += 1
	}
	similarityScore := calcSimilarityScore(leftArr, rightMap)
	fmt.Printf("(Part 2) The similarity score at the end of this process is %d\n", similarityScore)
}

func extractNumbersFromLine(line string) (int, int) {
	var currentNumber string
	var leftNum, rightNum int
	for _, char := range line {
		if isDigit(char) {
			currentNumber += string(char)
		}
		if char == ' ' && currentNumber != "" {
			num := strToInt(currentNumber)
			leftNum = num
			currentNumber = ""
		}
		if char == '\n' {
			num := strToInt(currentNumber)
			rightNum = num
		}
	}
	return leftNum, rightNum
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func strToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func insertMaintainingOrder(arr []int, num int) []int {
	if len(arr) == 0 || num < arr[0] {
		return insertAt(arr, 0, num)
	}
	if num > arr[len(arr)-1] {
		return insertAt(arr, len(arr), num)
	}
	insertPos := findInsertPos(arr, num)
	return insertAt(arr, insertPos, num)
}

func findInsertPos(arr []int, num int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := (l + r) / 2
		if arr[m] < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

func insertAt(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func calcTotalDistance(arr1, arr2 []int) int {
	totalDistance := 0
	for i, val1 := range arr1 {
		totalDistance += abs(val1 - arr2[i])
	}
	return totalDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcSimilarityScore(leftArr []int, rightMap map[int]int) int {
	similarityScore := 0
	for i := range leftArr {
		similarityScore += leftArr[i] * rightMap[leftArr[i]]
	}
	return similarityScore
}
