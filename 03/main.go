package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.ReadFile("input.txt")

	total := 0
	found := parse(file)
	for _, mul := range found {
		nums := parseNumbers(mul)
		num1, _ := strconv.Atoi(string(nums[0]))
		num2, _ := strconv.Atoi(string(nums[1]))
		total += num1 * num2
	}
	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Execution time: %s\n", time.Since(start))

	start = time.Now()
	totalWithEnablers := 0
	commandsWithEnablers := parseWithEnablers(file)
	i := 0
	do := true
	for i < len(commandsWithEnablers) {
		command := commandsWithEnablers[i]
		if string(command) == "do()" {
			do = true
			i++
			continue
		}
		if string(command) == "don't()" {
			do = false
			i++
			continue
		}
		if !do {
			i++
			continue
		}

		nums := parseNumbers(command)
		num1, _ := strconv.Atoi(string(nums[0]))
		num2, _ := strconv.Atoi(string(nums[1]))
		totalWithEnablers += num1 * num2
		i++
	}

	fmt.Printf("Part 2: %d\n", totalWithEnablers)
	fmt.Printf("Execution time: %s\n", time.Since(start))

}

func parse(file []byte) [][]byte {
	expression, _ := regexp.Compile(`(mul\(\d{1,3},\d{1,3}\))`)
	return expression.FindAll(file, -1)
}
func parseNumbers(mul []byte) [][]byte {
	expression, _ := regexp.Compile(`\d{1,3}`)
	return expression.FindAll(mul, -1)
}

func parseWithEnablers(file []byte) [][]byte {
	exp, _ := regexp.Compile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	return exp.FindAll(file, -1)
}
