package basex

import "testing"

func Test_pow(t *testing.T) {
	tests := []struct {
		base     int
		exponent int
		expected int
	}{
		{1, 0, 1},
		{2, 0, 1},
		{2, 2, 4},
		{3, 3, 27},
		{8, 2, 64},
		{8, 3, 512},
		{8, 4, 4096},
		{10, 4, 10000},
		{16, 4, 65536},
		{18, 3, 5832},
		{234, 2, 54756},
	}

	for _, test := range tests {
		if actual := pow(test.base, test.exponent); actual != test.expected {
			t.Errorf("pow(%d, %d) should give %d but received %d",
				test.base, test.exponent, test.expected, actual)
		}
	}
}
