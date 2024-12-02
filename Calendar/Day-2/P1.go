package day2

import (
	"strings"

	calendar "github.com/Harichandra-Prasath/AOC-24/Calendar"
)

func P1_Solve() int {

	data := calendar.Extract(INPUT_FILE)

	lines := strings.Split(data, "\n")

	count := len(lines)

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		if !isValidReport(numbers) {
			count -= 1
		}

	}

	return count

}
