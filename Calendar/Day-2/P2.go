package day2

import (
	"strings"

	calendar "github.com/Harichandra-Prasath/AOC-24/Calendar"
)

func P2_Solve() int {
	data := calendar.Extract(INPUT_FILE)

	lines := strings.Split(data, "\n")

	count := len(lines)

	// store the bad reports
	var bad_reports []string

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		if !isValidReport(numbers) {
			bad_reports = append(bad_reports, line)
			count -= 1
		}
	}

	for _, report := range bad_reports {
		numbers := strings.Split(report, " ")

		for i := range numbers {

			tmp := make([]string, len(numbers))
			copy(tmp, numbers)

			tmp = append(tmp[:i], tmp[i+1:]...)

			if isValidReport(tmp) {
				count += 1
				break
			}
		}

	}

	return count
}
