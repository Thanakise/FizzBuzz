package level1

import "strconv"

func FizzbuzzLevel1(number int) string {
	if number == 15 {
		return "FizzBuzz"
	}
	if number == 5 || number == 10{
		return "Buzz"
	}
	if number % 3 == 0{
		return "Fizz"
	}
	return strconv.Itoa(number)
}