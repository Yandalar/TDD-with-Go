package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectMessage(t, got, want, numbers)
	})
}

func assertCorrectMessage(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q. Given, %v", got, want, numbers)
	}
}
