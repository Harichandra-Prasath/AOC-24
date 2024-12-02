package day2

import "strconv"

const INPUT_FILE = `Calendar/Day-2/input.txt`

func isValidReport(numbers []string) bool {

	// Always a valid report
	if len(numbers) == 1 {
		return true
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
		return false
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
			return false
		}

	}

	return true

}
