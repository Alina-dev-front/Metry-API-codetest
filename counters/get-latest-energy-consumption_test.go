package counters_test

import (
	"testing"

	"github.com/Alina-dev-front/Metry-API-codetest/counters"
)

func TestGetLatestEnergyConsumption(t *testing.T) {
	scenarios := []struct {
		input    []int
		expected int
	}{
		{input: []int{5, 8, 2133, 0, 0, 0}, expected: 2133},
		{input: []int{17, 0}, expected: 17},
		{input: []int{6, 91, 811123, 5, 100}, expected: 100},
	}

	for _, scenario := range scenarios {
		got := counters.GetLatestEnergyConsumption(scenario.input)
		if got != scenario.expected {
			t.Errorf("Did not get expected result fot input '%v'. Got: '%v', expected: '%v'", scenario.input, got, scenario.expected)
		}
	}
}
