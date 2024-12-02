package day2

import (
	"strconv"
	"strings"

	calendar "github.com/Harichandra-Prasath/AOC-24/Calendar"
)

func P1_Solve() int {

	data := calendar.Extract(INPUT_FILE)

	lines := strings.Split(data, "\n")

	count := len(lines)

	for _, line := range lines {
		numbers := strings.Split(line, " ")

		// Always a valid report
		if len(numbers) == 1 {
			continue
		}

		var isAsc bool

		e1, _ := strconv.Atoi(numbers[0])
		e2, _ := strconv.Atoi(numbers[1])

		diff := e2 - e1
		if diff > 0 {
			isAsc = true
		} else if diff <= 0 {
			diff *= -1
			isAsc = false
		}

		if diff > 3 || diff < 1 {
			count -= 1
			continue
		}

		for i := 1; i < len(numbers)-1; i++ {
			n1, _ := strconv.Atoi(numbers[i])
			n2, _ := strconv.Atoi(numbers[i+1])

			var _diff int

			if isAsc {
				_diff = n2 - n1
			} else {
				_diff = n1 - n2
			}

			if _diff > 3 || _diff < 1 {
				count -= 1
				break
			}

		}

	}

	return count

}
