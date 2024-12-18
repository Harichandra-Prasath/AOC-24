package day1

import (
	"container/heap"
	"strconv"
	"strings"

	calendar "github.com/Harichandra-Prasath/AOC-24/Calendar"
)

func P1_Solve() int {

	data := calendar.Extract(INPUT_PATH)

	heap1 := &minHeap{}
	heap2 := &minHeap{}

	heap.Init(heap1)
	heap.Init(heap2)

	// Traverse each line
	for _, line := range strings.Split(data, "\n") {

		ids := strings.Split(line, "   ")

		id_1, _ := strconv.Atoi(ids[0])
		id_2, _ := strconv.Atoi(ids[1])

		heap.Push(heap1, id_1)
		heap.Push(heap2, id_2)

	}

	// Traverse the heap and get the sum
	n := heap1.Len()
	sum := 0

	for i := 0; i < n; i++ {
		left := heap.Pop(heap1).(int)
		right := heap.Pop(heap2).(int)

		curr := right - left

		if curr < 0 {
			curr *= -1
		}

		sum += curr

	}

	return sum

}
