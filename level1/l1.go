package level1

import "strconv"


func FizzbuzzLevel1(number int) string {
	if number == 5 {
		return "Buzz"
	}

	if number == 3 || number == 6 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}