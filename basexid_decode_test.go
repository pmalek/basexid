package basex

import "testing"

// Decode

func Test_Decode(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"0", ""},
		{"1", "1"},
		{"120", "1w"},
		{"198237", "pZN"},
		{"123987123176", "2BKvKfI"},
	}

	for _, test := range tests {
		if actual := Encode(test.value); actual != test.expected {
			t.Errorf("Encode(%q) should give %q but received %q",
				test.value, test.expected, actual)
		}
	}
}
