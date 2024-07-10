package level1

import "strconv"


func FizzbuzzLevel1(number int) string {
	text := ""
	flag := false

	if number % 3 == 0 {
		text = "Fizz"
		flag = true
	} 
	
	if number == 5 {
		text = "Buzz"
		flag = true
	}
	
	if !flag {
		text = strconv.Itoa(number)
	}
	return text
}