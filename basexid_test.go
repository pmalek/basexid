package basex

import "testing"

func test(exp, act string, t *testing.T) {
	if exp != act {
		t.Errorf("Expected: '%s', got '%s'", exp, act)
	}
}

func test_int(exp, act int, t *testing.T) {
	if exp != act {
		t.Errorf("Expected: '%d', got '%d'", exp, act)
	}
}
