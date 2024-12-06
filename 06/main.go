package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var guardMap [][]string

func main() {
	start := time.Now()
	guardMap = readFile("input.txt")
	i, j, direction := findStart()
	walkGuard(i, j, direction)
	printGuardMap()
	paths := countWalkedPath()
	fmt.Printf("Number of paths: %d\n", paths)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}
func readFile(fileName string) [][]string {
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cells := make([]string, len(line))
		guardMap = append(guardMap, cells)
		for i, c := range line {
			cells[i] = string(c)
		}
	}
	return guardMap
}

func walkGuard(i, j int, direction string) {
	guardMap[i][j] = "X"
	if direction == "^" && i == 0 {
		return
	} else if direction == ">" && j == len(guardMap[0])-1 {
		return
	} else if direction == "v" && i == len(guardMap)-1 {
		return
	} else if direction == "<" && j == 0 {
		return
	}

	var nextCell string

	if direction == "^" {
		nextCell = guardMap[i-1][j]
	} else if direction == ">" {
		nextCell = guardMap[i][j+1]
	} else if direction == "v" {
		nextCell = guardMap[i+1][j]
	} else {
		nextCell = guardMap[i][j-1]
	}

	if nextCell == "#" {
		direction = changeDirection(direction)
	}
	if direction == "^" {
		guardMap[i-1][j] = "^"
		walkGuard(i-1, j, direction)
	} else if direction == "v" {
		guardMap[i+1][j] = "v"
		walkGuard(i+1, j, direction)
	} else if direction == "<" {
		guardMap[i][j-1] = "<"
		walkGuard(i, j-1, direction)
	} else {
		guardMap[i][j+1] = ">"
		walkGuard(i, j+1, direction)
	}
}

func changeDirection(direction string) string {
	if direction == "^" {
		direction = ">"
	} else if direction == ">" {
		direction = "v"
	} else if direction == "v" {
		direction = "<"
	} else {
		direction = "^"
	}
	return direction
}

func findStart() (int, int, string) {
	for i, row := range guardMap {
		for j, cell := range row {
			if cell == "^" || cell == "v" || cell == "<" || cell == ">" {
				return i, j, cell
			}
		}
	}
	return 0, 0, ""
}

func printGuardMap() {
	for _, row := range guardMap {
		fmt.Println(row)
	}
	fmt.Println()
}

func countWalkedPath() int {
	paths := 0
	for _, row := range guardMap {
		for _, cell := range row {
			if cell == "X" {
				paths++
			}
		}
	}
	return paths
}
