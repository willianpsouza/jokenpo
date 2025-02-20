package randomize

import (
	"testing"
)

func TestRealRandomSeed(t *testing.T) {
	result := PseudoPrimeGenerator()
	if result <= 0 {
		t.Errorf("Expected PseudoPrimeGenerator() to return a positive number, got %d", result)
	}
}

func TestMultipleCalls(t *testing.T) {
	first := PseudoPrimeGenerator()
	second := PseudoPrimeGenerator()

	if first == second {
		t.Errorf("Expected different values from multiple calls, but got identical values: %d and %d", first, second)
	}
}

func TestRandomSeedRange(t *testing.T) {
	const minExpected = 1000000
	const maxExpected = 1000000000000

	result := PseudoPrimeGenerator()

	if result < minExpected || result > maxExpected {
		t.Errorf("PseudoPrimeGenerator() returned %d, which is out of expected range [%d, %d]", result, minExpected, maxExpected)
	}
}
