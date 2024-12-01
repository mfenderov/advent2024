package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"github.com/pkg/errors"
	"math"
	"os"
	"strconv"
	"strings"
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

func main() {

	heap1 := &IntHeap{}
	heap2 := &IntHeap{}
	heap.Init(heap1)
	heap.Init(heap2)

	err := readIntoHeaps(heap1, heap2)
	if err != nil {
		panic(err)
	}

	total := 0

	for heap1.Len() > 0 && heap2.Len() > 0 {
		num1 := heap.Pop(heap1).(int)
		num2 := heap.Pop(heap2).(int)

		total += int(math.Abs((float64(num1 - num2))))

	}

	fmt.Println(total)
}

func readIntoHeaps(heap1 *IntHeap, heap2 *IntHeap) error {
	file, err := os.Open("input.txt")
	if err != nil {
		return errors.Wrap(err, "error opening file")
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
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "error reading file")
	}

	return nil
}
