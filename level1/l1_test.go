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
