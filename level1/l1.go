package level1

import "strconv"


func ConvertIntToString(number int) string{
	return strconv.Itoa(number)
}

func FizzbuzzLevel1(number int) string {
	if number == 3 {
		return "Fizz"
	}
	return ConvertIntToString(number)
}