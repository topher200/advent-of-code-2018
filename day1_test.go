package main

import "testing"

func TestPositiveOnes(t *testing.T) {
	positiveOnes := []string{
		"+1",
		"+1",
		"+1",
		"+1",
	}
	expectedValue := 4
	calculatedValue := ApplyListOfValues(positiveOnes)
	if expectedValue != calculatedValue {
		t.Errorf("Expected %d, got %d", expectedValue, calculatedValue)
	}
}

func TestNegativeOnes(t *testing.T) {
	inputList := []string{
		"-1",
		"-1",
		"-1",
		"-1",
	}
	expectedValue := -4
	calculatedValue := ApplyListOfValues(inputList)
	if expectedValue != calculatedValue {
		t.Errorf("Expected %d, got %d", expectedValue, calculatedValue)
	}
}
