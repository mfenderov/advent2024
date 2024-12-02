package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Change to `>` for max-heap
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type InputData struct {
	LeftHeap  *IntHeap
	RightHeap *IntHeap
	LeftSlice []int
	RightMap  map[int]int
}

func main() {

	inputData, err := readIntoHeaps()

	if err != nil {
		panic(err)
	}

	total := 0

	leftHeap := inputData.LeftHeap
	rightHeap := inputData.RightHeap
	for leftHeap.Len() > 0 && rightHeap.Len() > 0 {
		num1 := heap.Pop(leftHeap).(int)
		num2 := heap.Pop(rightHeap).(int)

		total += int(math.Abs(float64(num1 - num2)))
	}

	fmt.Printf("total distance: %d \n", total)

	leftSlice := inputData.LeftSlice
	rightMap := inputData.RightMap
	similarityScore := 0
	for _, left := range leftSlice {
		frequency := rightMap[left]
		similarityScore += left * frequency
	}

	fmt.Printf("similarity score: %d \n", similarityScore)

}

func readIntoHeaps() (InputData, error) {

	heap1 := &IntHeap{}
	heap2 := &IntHeap{}
	heap.Init(heap1)
	heap.Init(heap2)

	var slice1 []int
	map2 := make(map[int]int)

	file, err := os.Open("input.txt")
	if err != nil {
		return InputData{}, errors.Wrap(err, "error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing line:", line)
			continue
		}

		heap.Push(heap1, num1)
		heap.Push(heap2, num2)

		slice1 = append(slice1, num1)
		map2[num2]++
	}

	if err := scanner.Err(); err != nil {
		return InputData{}, errors.Wrap(err, "error reading file")
	}

	return InputData{
		LeftHeap:  heap1,
		RightHeap: heap2,
		LeftSlice: slice1,
		RightMap:  map2,
	}, nil
}
