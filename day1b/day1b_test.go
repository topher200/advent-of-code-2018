package main

import "testing"

func TestLoop(t *testing.T) {
	input := []string{
		"+1",
		"+1",
		"+1",
		"-4",
	}
	expectedValue := 1
	calculatedValue := FindFirstRepeatFrequency(input)
	if expectedValue != calculatedValue {
		t.Errorf("Expected %d, got %d", expectedValue, calculatedValue)
	}
}
