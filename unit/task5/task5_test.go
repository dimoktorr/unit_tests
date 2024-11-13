package task5_test

import "testing"

func Test_sum(t *testing.T) {
	//Arrange
	a := 5
	b := 5
	want := 10

	//Act
	got := sum(a, b)

	//Assert
	if got != want {
		t.Errorf("sum(%d, %d) = %d; want %d", a, b, got, want)
	}
}
