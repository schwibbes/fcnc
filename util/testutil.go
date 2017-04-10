package util

import (
	"testing"
)

func AssertTrue(result bool, t *testing.T) {
	if !result {
		t.Fatal("expect true, was false")
	}
}
func AssertFalse(result bool, t *testing.T) {
	if result {
		t.Fatal("expect false, was true")
	}
}

func AssertEqual(a, b interface{}, t *testing.T) {
	if a != b {
		t.Fatalf("params were not equal, %v != %v", a, b)
	}
}
