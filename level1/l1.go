package level1

import "strconv"


func ConvertIntToString(number int) string{
	return strconv.Itoa(number)
}

func isDivideBy5(number int) bool{
	return number % 5 == 0
}

func isDivideBy3(number int) bool{
	return number % 3 == 0
}

func FizzbuzzLevel1(number int) string {

	if isDivideBy3(number) && isDivideBy5(number){
		return "FizzBuzz"
	}

	if isDivideBy5(number){
		return "Buzz"
	}

	if isDivideBy3(number){
		return "Fizz"
	}
	return ConvertIntToString(number)
}