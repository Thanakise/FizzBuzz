package level1

import "testing"

func TestFizzBuzzLevel1(t *testing.T){
	input := 1
	got := FizzbuzzLevel1(input)
	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel2(t *testing.T){
	input := 2
	got := FizzbuzzLevel1(input)
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
func TestFizzBuzzLevel4(t *testing.T){
	input := 4
	got := FizzbuzzLevel1(input)
	want := "4"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel5(t *testing.T){
	input := 5
	got := FizzbuzzLevel1(input)
	want := "Buzz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel6(t *testing.T){
	input := 6
	got := FizzbuzzLevel1(input)
	want := "Fizz"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
func TestFizzBuzzLevel7(t *testing.T){
	input := 7
	got := FizzbuzzLevel1(input)
	want := "7"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
