package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	totalValid, totalInvalid := readFile("input.txt")
	fmt.Printf("valid: %d", totalValid)
	fmt.Println()
	fmt.Printf("fixed: %d", totalInvalid)
	fmt.Println()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func readFile(fileName string) (int, int) {
	pageMap, toValidate := parseInput(fileName)

	totalValid := 0
	totalInvalid := 0
	for _, v := range toValidate {
		isValid := true
		for i := len(v) - 1; i >= 0; i-- {
			last := v[i]
			after := pageMap[last]
			if len(after) == 0 {
				continue
			}
			if intersection(after, i, v) {
				isValid = false
				break
			}
		}
		if isValid {
			mid, _ := strconv.Atoi(v[len(v)/2])
			totalValid += mid
		} else {
			sort.Slice(v, comparator(v, pageMap))
			mid, _ := strconv.Atoi(v[len(v)/2])
			totalInvalid += mid
		}
	}
	return totalValid, totalInvalid
}

func comparator(v []string, pageMap map[string][]string) func(i int, j int) bool {
	return func(i, j int) bool {
		vi := v[i]
		vj := v[j]

		afterI := pageMap[vi]
		if len(afterI) != 0 {
			for _, a := range afterI {
				if a == vj {
					return true
				}
			}
		}
		afterJ := pageMap[vj]
		if len(afterJ) != 0 {
			for _, a := range afterJ {
				if a == vi {
					return false
				}
			}
		}
		return true
	}
}

func intersection(after []string, i int, v []string) bool {
	for _, a := range after {
		for j := 0; j < i; j++ {
			if a == v[j] {
				return true
			}
		}
	}
	return false
}

func parseInput(fileName string) (map[string][]string, [][]string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pageMap := make(map[string][]string)
	toValidate := make([][]string, 0)
	firstHalfOfTheInput := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstHalfOfTheInput = false
			continue
		}
		if firstHalfOfTheInput {
			pages := strings.Split(line, "|")
			before := pages[0]
			after := pages[1]
			pageMap[before] = append(pageMap[before], after)
		}
		if !firstHalfOfTheInput {
			toValidate = append(toValidate, strings.Split(line, ","))
		}
	}
	return pageMap, toValidate
}
