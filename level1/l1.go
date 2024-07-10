package level1

import "strconv"

func FizzbuzzLevel1(number int) string {
	if number == 3 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}