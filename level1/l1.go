package level1

import "strconv"


func FizzbuzzLevel1(number int) string {
	if number % 5 == 0{
		return "Buzz"
	}

	if number % 3 == 0{
		return "Fizz"
	}
	return strconv.Itoa(number)
}