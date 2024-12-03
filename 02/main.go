package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func main() {

	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	safeReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if validate(parts, 0) {
			safeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(errors.Wrap(err, "error reading file"))
	}

	fmt.Printf("Safe reports: %d\n", safeReports)
	fmt.Printf("Execution time: %s\n", time.Since(start))

}

func validate(parts []string, errors int) bool {
	isDescending := false
	for i, p := range parts {
		if i == len(parts)-1 {
			break
		}
		part, _ := strconv.Atoi(p)
		nextPart, _ := strconv.Atoi(parts[i+1])

		if i == 0 && part > nextPart {
			isDescending = true
		}
		abs := math.Abs(float64(part - nextPart))
		if (abs < 1 || abs > 3) ||
			(isDescending && part < nextPart) ||
			(!isDescending && part > nextPart) {
			if errors > 0 {
				return false
			}

			withoutLeft := remove(parts, i)
			withoutRight := remove(parts, i+1)

			withoutFirst := remove(parts, 0)
			leftIsValid := validate(withoutLeft, errors+1)
			rightIsValid := validate(withoutRight, errors+1)
			firstIsValid := validate(withoutFirst, errors+1)
			if leftIsValid || rightIsValid || firstIsValid {
				return true
			}
			return false
		}
	}
	return true
}

func remove(slice []string, s int) []string {
	result := append([]string{}, slice[:s]...)
	return append(result, slice[s+1:]...)
}
