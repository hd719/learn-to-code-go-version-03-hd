package main

import "testing"

// To test functions in Go, we need to create a file with the name ending in _test.go and then create a function with the name TestXxx (where Xxx does not start with a lowercase letter) in that file.
// The test function expects a single parameter t *testing.T, which is a type defined in the testing package.
// *testing.T (pointer to testing) is a type that has methods we can call to indicate whether our test has passed or failed.
// Run go test to run the test
func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
