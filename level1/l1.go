package level1

import "strconv"


func FizzbuzzLevel1(number int) string {
	text := ""


	if number == 3 {
		text = "Fizz"
	} else {
		text = strconv.Itoa(number)
	}
	return text
}