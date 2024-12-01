package day1

import (
	"strconv"
	"strings"

	calendar "github.com/Harichandra-Prasath/AOC-24/Calendar"
)

func P2_Solve() int {

	data := calendar.Extract(INPUT_PATH)

	store := make(map[string]int)

	// To store the id1
	var IDs []string

	for _, line := range strings.Split(data, "\n") {
		ids := strings.Split(line, "   ")

		id_1 := ids[0]
		id_2 := ids[1]

		IDs = append(IDs, id_1)
		store[id_2] += 1
	}

	similarity := 0

	// Traverse the store
	for _, _id_1 := range IDs {

		freq := store[_id_1]

		id_1, _ := strconv.Atoi(_id_1)

		similarity += (id_1 * freq)
	}

	return similarity
}
