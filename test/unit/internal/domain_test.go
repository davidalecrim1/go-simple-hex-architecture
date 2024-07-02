package internal_test

import (
	internal "go_simple_hex_architecture/internal"
	"testing"
)

func TestNewTax(t *testing.T) {
	id := "tax001"
	tax := internal.NewTax(id)

	if tax.ID != id {
		t.Errorf("Expected tax ID to be %s, but got %s", id, tax.ID)
	}

	if tax.Tax != 0 {
		t.Errorf("Expected initial tax value to be 0, but got %f", tax.Tax)
	}
}

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name     string
		price    float64
		rate     float64
		expected float64
	}{
		{"Positive values", 100.0, 0.1, 110.0},
		{"Zero price", 0.0, 0.2, 0.0},
		{"Zero rate", 200.0, 0.0, 200.0},
		{"Negative price", -50.0, 0.15, -57.5},
		{"Negative rate", 75.0, -0.05, 71.25},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tax := internal.NewTax("test")
			tax.Calculate(tc.price, tc.rate)

			if tax.Tax != tc.expected {
				t.Errorf("Expected tax value to be %f, but got %f", tc.expected, tax.Tax)
			}
		})
	}
}
