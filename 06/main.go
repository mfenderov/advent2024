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
	paths := countWalkedPath()
	guardMap[i][j] = direction
	searchForLoops(i, j, direction)
	fmt.Printf("Number of xpaths: %d\n", paths)
	fmt.Printf("Number of loops: %d\n", loopsFounds)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func searchForLoops(i, j int, direction string) {
	for a, row := range guardMap {
		for b, cell := range row {
			ab := fmt.Sprintf("%d,%d", a, b)
			ij := fmt.Sprintf("%d,%d", i, j)
			if ab != ij && cell == "X" {
				tmpMap := make([][]string, len(guardMap))
				for z, zrow := range guardMap {
					tmpMap[z] = make([]string, len(zrow))
					copy(tmpMap[z], zrow)
				}
				guardMap[a][b] = "#"
				if walkGuard(i, j, direction) {
					loopsFounds++
				}

				for x, xrow := range tmpMap {
					guardMap[x] = make([]string, len(xrow))
					copy(guardMap[x], xrow)
				}
				visited = make(map[string]int)
			}
		}
	}
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

var visited = make(map[string]int)
var loopsFounds = 0

func walkGuard(i, j int, direction string) bool {
	if amILooping(i, j, direction) {
		return true
	}
	guardMap[i][j] = "X"
	visited[fmt.Sprintf("%d,%d,%s", i, j, direction)]++
	if direction == "^" && i == 0 {
		return false
	} else if direction == ">" && j == len(guardMap[0])-1 {
		return false
	} else if direction == "v" && i == len(guardMap)-1 {
		return false
	} else if direction == "<" && j == 0 {
		return false
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
		return walkGuard(i-1, j, direction)
	} else if direction == "v" {
		guardMap[i+1][j] = "v"
		return walkGuard(i+1, j, direction)
	} else if direction == "<" {
		guardMap[i][j-1] = "<"
		return walkGuard(i, j-1, direction)
	} else {
		guardMap[i][j+1] = ">"
		return walkGuard(i, j+1, direction)
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

func amILooping(i, j int, direction string) bool {
	if visited[fmt.Sprintf("%d,%d,%s", i, j, direction)] > 2 {
		return true
	}
	return false
}
