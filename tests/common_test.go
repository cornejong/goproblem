package tests

import (
	"fmt"
	"testing"
)

func AssertTest(t *testing.T, value any, expect any, message string) {
	defer t.Cleanup(func() {
		value = nil
		expect = nil
		message = ""
	})

	if value != expect {
		fmt.Printf("\n----------------------------------------\n✖ Assertion failed: %s\nExpected '%v'. Got: '%v'\n----------------------------------------\n\n", message, expect, value)
		t.FailNow()
	}
}
