package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Prisma struct {
	l int
	w int
	h int
}

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	var wrappingPapper, ribbonFeet int
	line, _, err := reader.ReadLine()
	for err != io.EOF {
		prisma := createPrisma(line)
		wrappingPapper += calcPresentArea(prisma)
		ribbonFeet += calcRibbonArea(prisma)
		line, _, err = reader.ReadLine()
	}
	fmt.Println("(Part 1) Total wrapping paper needed: ", wrappingPapper)
	fmt.Println("(Part 2) Total paper needed for ribbons: ", ribbonFeet)
}

func calcPresentArea(p Prisma) int {
	side1 := p.l * p.w
	side2 := p.w * p.h
	side3 := p.h * p.l
	area := 2 * (side1 + side2 + side3)
	return area + min(side1, side2, side3)
}

func createPrisma(line []byte) Prisma {
	splitLine := strings.Split(string(line), "x")
	l, _ := strconv.Atoi(splitLine[0])
	w, _ := strconv.Atoi(splitLine[1])
	h, _ := strconv.Atoi(splitLine[2])
	return Prisma{l, w, h}
}

func calcRibbonArea(p Prisma) int {
	sidesSum := p.l + p.w + p.h
	smallestSides := sidesSum - max(p.l, p.w, p.h)
	smallestPerimiter := 2 * smallestSides
	cubic := p.l * p.w * p.h
	return smallestPerimiter + cubic
}
