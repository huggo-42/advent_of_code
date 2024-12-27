package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
}

func part1() {
	ans := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		ans += parseInstructions(line)
	}
	fmt.Printf("THE END: %d\n", ans)
}

func parseInstructions(line string) int {
	var result int
	pos := 0
	for pos < len(line) {
		if isMultiplicationOperation(line[pos:]) {
			pos += 3
			left, indexToResume, err := parseNumberUntil(line[pos+1:], ',')
			if err != nil {
				continue
			}
			pos += indexToResume + 2
			right, indexToResume, err := parseNumberUntil(line[pos:], ')')
			if err != nil {
				continue
			}
			pos += indexToResume + 1
			result += left * right
			continue
		}
		pos++
	}
	return result
}

func parseNumberUntil(s string, delim byte) (int, int, error) {
	var numStr string
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			numStr += string(s[i])
			continue
		}
		if s[i] == delim && numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, 0, fmt.Errorf("invalid number %s: %w", numStr, err)
			}
			return num, i, nil
		}
		return 0, 0, fmt.Errorf("unexpected character: %c", s[i])
	}
	return 0, 0, fmt.Errorf("delimiter %c not found", delim)
}

func isMultiplicationOperation(s string) bool {
	return len(s) >= 4 && s[:4] == "mul("
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
