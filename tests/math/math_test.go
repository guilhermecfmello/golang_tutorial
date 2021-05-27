package math

import "testing"

const defaultError = "Expected value: %v, but result was %v."

func TestAverage(t *testing.T) {
	expectedValue := 15.00
	value := Average(10, 20)

	if value != expectedValue {
		t.Errorf(defaultError, expectedValue, value)
	}
}
