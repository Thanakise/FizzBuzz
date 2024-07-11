package level1

import "testing"

func TestFizzBuzzLevel1(t *testing.T){
	input := 1
	got := ConvertIntToString(input)
	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel2(t *testing.T){
	input := 2
	got := ConvertIntToString(input)
	want := "2"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel3(t *testing.T){
	input := 3
	got := FizzbuzzLevel1(input)
	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
// func TestFizzBuzzLevel4(t *testing.T){
// 	input := 4
// 	got := FizzbuzzLevel1(input)
// 	want := "4"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel5(t *testing.T){
// 	input := 5
// 	got := FizzbuzzLevel1(input)
// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel6(t *testing.T){
// 	input := 6
// 	got := FizzbuzzLevel1(input)
// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel7(t *testing.T){
// 	input := 7
// 	got := FizzbuzzLevel1(input)
// 	want := "7"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel8(t *testing.T){
// 	input := 8
// 	got := FizzbuzzLevel1(input)
// 	want := "8"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel9(t *testing.T){
// 	input := 9
// 	got := FizzbuzzLevel1(input)
// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel10(t *testing.T){
// 	input := 10
// 	got := FizzbuzzLevel1(input)
// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel11(t *testing.T){
// 	input := 11
// 	got := FizzbuzzLevel1(input)
// 	want := "11"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel12(t *testing.T){
// 	input := 12
// 	got := FizzbuzzLevel1(input)
// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel13(t *testing.T){
// 	input := 13
// 	got := FizzbuzzLevel1(input)
// 	want := "13"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel14(t *testing.T){
// 	input := 14
// 	got := FizzbuzzLevel1(input)
// 	want := "14"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
// func TestFizzBuzzLevel15(t *testing.T){
// 	input := 15
// 	got := FizzbuzzLevel1(input)
// 	want := "FizzBuzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
