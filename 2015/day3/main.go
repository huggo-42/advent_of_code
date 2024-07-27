package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x int
	y int
}

const (
	NORTH = 94
	SOUTH = 118
	EAST  = 62
	WEST  = 60
)

func main() {
	part1()
	part2()
}

func part1() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	housesVisited := 1
	currentPos := Pos{0, 0}
	visitedHouses := []Pos{currentPos}
	for {
		direction, err := reader.ReadByte()
		if err != nil {
			break
		}
		currentPos = moveTo(currentPos, direction)
		alreadyVisited := hasVisitedPos(&visitedHouses, &currentPos)
		if !alreadyVisited {
			housesVisited += 1
			visitedHouses = append(visitedHouses, currentPos)
		}
	}
	fmt.Println("(Part 1) Houses that received at least one present:", housesVisited)
}

func part2() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	housesVisitedByBoth := 1
	santaPos, roboPos := Pos{0, 0}, Pos{0, 0}
	visitedHousesPos := []Pos{santaPos, roboPos}
	santaMove := true
	for {
		var posToCheck Pos
		direction, err := reader.ReadByte()
		if err != nil {
			break
		}
		if santaMove {
			santaPos = moveTo(santaPos, direction)
			posToCheck = santaPos
		} else {
			roboPos = moveTo(roboPos, direction)
			posToCheck = roboPos
		}
		alreadyVisited := hasVisitedPos(&visitedHousesPos, &posToCheck)
		if !alreadyVisited {
			housesVisitedByBoth += 1
			visitedHousesPos = append(visitedHousesPos, posToCheck)
		}
		santaMove = !santaMove
	}
	fmt.Println("(Part 2) Houses visited by Santa and Robo-Santa: ", housesVisitedByBoth)
}

func moveTo(pos Pos, direction byte) Pos {
	switch direction {
	case NORTH:
		pos.y += 1
	case SOUTH:
		pos.y -= 1
	case EAST:
		pos.x += 1
	case WEST:
		pos.x -= 1
	}
	return pos
}

func hasVisitedPos(visitedHousesPos *[]Pos, pos *Pos) bool {
	hasVisited := false
	for _, visitedPos := range *visitedHousesPos {
		if visitedPos == *pos {
			hasVisited = true
		}
	}
	return hasVisited
}
