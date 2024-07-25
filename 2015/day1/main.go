package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FLOOR_UP   = 40
	FLOOR_DOWN = 41
)

func main() {
	file, _ := os.OpenFile("input.txt", 0, 0x0)
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	floor := part1(line)
	position := part2(line)
	fmt.Println("(Part 1) The instructions takes santa to floor ", floor)
	fmt.Println("(Part 2) Santa enters the basement at position ", position)
}

func part1(line string) int {
	var floor int
	for _, char := range line {
		switch char {
		case FLOOR_UP:
			floor++
		case FLOOR_DOWN:
			floor--
		}
	}
	return floor
}

func part2(line string) int {
	var floor, position int
	for i, char := range line {
		switch char {
		case FLOOR_UP:
			floor++
		case FLOOR_DOWN:
			floor--
		}
		if floor == -1 {
			position = i + 1
			break
		}
	}
	return position
}
