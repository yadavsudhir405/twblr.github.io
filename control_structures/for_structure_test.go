package control_structures

import "testing"

func TestMapInts(t *testing.T) {

	square := func(arg1 int32) int32 {
		return arg1 * arg1
	}

	res := mapInts(square, []int32{1, 2, 3})
	expectedResult := []int32{1, 4, 9}

	if len(res) < 1 {
		t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
	}

	for i := range res {
		if res[i] != expectedResult[i] {
			t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
		}
	}
}

func TestFilterInts(t *testing.T) {

	greaterThanTwo := func(arg1 int32) bool {
		return arg1 > 2
	}

	res := filterInts(greaterThanTwo, []int32{1, 2, 3})
	expectedResult := []int32{3}

	if len(res) < 1 {
		t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
	}

	for i := range res {
		if res[i] != expectedResult[i] {
			t.Errorf("Expected the result to be %v, but got %v", expectedResult, res)
		}
	}
}
