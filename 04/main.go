package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	xmases, mases := findAllXmasInFile("input.txt")
	fmt.Printf("XMAS found: %d\n", xmases)
	fmt.Printf("MAS found: %d\n", mases)
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

var matrix [][]string

func findAllXmasInFile(fileName string) (int, int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	matrix = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []string{})
		for _, v := range line {
			matrix[row] = append(matrix[row], string(v))
		}
		row++
	}

	return findAllXMASes(), findAllMASes()
}

func findAllXMASes() int {
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				result += exploreAround(i, j)
			}
		}
	}
	return result
}

func exploreAround(i int, j int) int {
	result := 0

	xmas := ""
	if i+3 < len(matrix) {
		X := matrix[i][j]
		M := matrix[i+1][j]
		A := matrix[i+2][j]
		S := matrix[i+3][j]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if i-3 >= 0 {
		X := matrix[i][j]
		M := matrix[i-1][j]
		A := matrix[i-2][j]
		S := matrix[i-3][j]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if j+3 < len(matrix[i]) {
		X := matrix[i][j]
		M := matrix[i][j+1]
		A := matrix[i][j+2]
		S := matrix[i][j+3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if j-3 >= 0 {
		X := matrix[i][j]
		M := matrix[i][j-1]
		A := matrix[i][j-2]
		S := matrix[i][j-3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if i+3 < len(matrix) && j+3 < len(matrix[i]) {
		X := matrix[i][j]
		M := matrix[i+1][j+1]
		A := matrix[i+2][j+2]
		S := matrix[i+3][j+3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if i-3 >= 0 && j-3 >= 0 {
		X := matrix[i][j]
		M := matrix[i-1][j-1]
		A := matrix[i-2][j-2]
		S := matrix[i-3][j-3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if i+3 < len(matrix) && j-3 >= 0 {
		X := matrix[i][j]
		M := matrix[i+1][j-1]
		A := matrix[i+2][j-2]
		S := matrix[i+3][j-3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}
	if i-3 >= 0 && j+3 < len(matrix[i]) {
		X := matrix[i][j]
		M := matrix[i-1][j+1]
		A := matrix[i-2][j+2]
		S := matrix[i-3][j+3]

		xmas = X + M + A + S
		result += validateXmas(xmas)
	}

	return result
}

func validateXmas(xmas string) int {
	if xmas == "XMAS" || xmas == "SAMX" {
		return 1
	}
	return 0
}

func findAllMASes() int {
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "A" {
				result += exploreAroundMAS(i, j)
			}
		}
	}
	return result
}

func exploreAroundMAS(i int, j int) int {
	mas := ""
	A := matrix[i][j]
	if i+1 < len(matrix) && i-1 >= 0 && j+1 < len(matrix[i]) && j-1 >= 0 &&
		i+1 < len(matrix) && i-1 >= 0 && j+1 < len(matrix[i]) && j-1 >= 0 {
		M := matrix[i-1][j-1]
		S := matrix[i+1][j+1]
		mas = M + A + S
		if !validateMAS(mas) {
			return 0
		}
		M = matrix[i-1][j+1]
		S = matrix[i+1][j-1]
		mas = M + A + S
		if !validateMAS(mas) {
			return 0
		}
		return 1
	}
	return 0
}

func validateMAS(mas string) bool {
	if mas == "MAS" || mas == "SAM" {
		return true
	}
	return false
}
