package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxRed = 12
const maxBlue = 14
const maxGreen = 13

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var possibleGamesIdsSum, minSetOfCubes int
	for fileScanner.Scan() {
		possibleGamesIdsSum += validateGamesInLine(fileScanner.Text())
		minSetOfCubes += powerOfMinSetOfCubes(fileScanner.Text())
	}
	readFile.Close()
	fmt.Println("(Part 1) Possible games ids sum: ", possibleGamesIdsSum)
	fmt.Println("(Part 2) Sum of the power of the minimum set of cubes: ", minSetOfCubes)
}

// returns gameId if game is possible and 0 if game is not possible
func validateGamesInLine(line string) int {
	var s = strings.Split(line, ":")
	var gameId, _ = strconv.Atoi(strings.Split(s[0], " ")[1])
	var plays = strings.Split(s[1], ";")
	var possiblePlays int
	for _, play := range plays {
		var red, green, blue int
		var num string
		for _, char := range play {
			if isNumber(char) {
				num += string(char)
			} else {
				var intNum, _ = strconv.Atoi(num)
				switch char {
				case 'r':
					red += intNum
				case 'g':
					green += intNum
				case 'b':
					blue += intNum
				default:
					continue
				}
				num = ""
			}
		}
		if isPlayPossible(red, green, blue) {
			possiblePlays += 1
		}
	}
	if possiblePlays == len(plays) {
		return gameId
	}
	return 0
}

// returns the power of the minimun set of cubes needed for the game
func powerOfMinSetOfCubes(line string) int {
	var s = strings.Split(line, ":")
	var plays = strings.Split(s[1], ";")
	var minRed, minGreen, minBlue int
	for _, play := range plays {
		var red, green, blue int
		var num string
		for _, char := range play {
			if isNumber(char) {
				num += string(char)
			} else {
				var intNum, _ = strconv.Atoi(num)
				switch char {
				case 'r':
					red += intNum
				case 'g':
					green += intNum
				case 'b':
					blue += intNum
				default:
					continue
				}
				num = ""
			}
		}
		if red > minRed {
			minRed = red
		}
		if green > minGreen {
			minGreen = green
		}
		if blue > minBlue {
			minBlue = blue
		}
	}
	return minRed * minGreen * minBlue
}

func isPlayPossible(red int, green int, blue int) bool {
	return red <= maxRed && blue <= maxBlue && green <= maxGreen
}
func isNumber(n int32) bool {
	return n > 47 && n < 58
}
