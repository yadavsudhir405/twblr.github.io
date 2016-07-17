package control_structures

import "testing"

var calcTests = []struct {
	op     string
	arg1   float32
	arg2   float32
	output float32
}{
	{add, 1, 2, 3},
	{subtract, 4, 8, -4},
	{multiply, 4, 0, 0},
	{divide, 2, 4, 0.5},
}

func TestCalculatorSum(t *testing.T) {

	for _, tt := range calcTests {
		res := calculate(tt.op, tt.arg1, tt.arg2)
		if res != tt.output {
			t.Errorf("Expected the result to be %f,  but instead got %f", tt.output, res)
		}
	}
}
