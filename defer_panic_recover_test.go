package main

import "testing"

func TestDefer(t *testing.T) {
	res := allocateResource()
	if res.closed != true {
		t.Errorf("Expected the resource to be closed")
	}
}

func TestPanic(t *testing.T) {
	err := handlePanicOpen()
	if err.Error() != "resource panicked" {
		t.Errorf("Did not recover from panic")
	}
}
