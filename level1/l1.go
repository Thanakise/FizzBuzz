package level1


func FizzbuzzLevel1(number int) string {
	text := ""


	if number == 3 {
		text = "Fizz"
	}

	if number == 2 {
		text = "2"
	}

	if number == 1 {
		text = "1"
	}
	return text
}