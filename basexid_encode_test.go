package basex

import "testing"

func Test_Encode(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"-57", ""},
		{"0", ""},
		{"1", "1"},
		{"120", "1w"},
		{"198237", "pZN"},
		{"1871176", "7qmG"},
		{"12812317", "rl4H"},
		{"123812176", "8NVCa"},
		{"123987123176", "2BKvKfI"},
		{"1239871231761", "LpNFKev"},
		{"62398712317612121", "4bmmJzNcjZ"},
		{"623987123176121212", "k5rnDtoFLg"},
		{"923987123176121212", "16FrzSRUP4W"},
	}

	for _, test := range tests {
		if actual := Encode(test.value); actual != test.expected {
			t.Errorf("Encode(%q) should give %q but received %q",
				test.value, test.expected, actual)
		}
	}
}

// Encodei

func Test_Encodei(t *testing.T) {
	tests := []struct {
		value    int
		expected string
	}{
		{-57, ""},
		{0, ""},
		{1, "1"},
		{120, "w1"},
		{198237, "NZp"},
		{1871176, "Gmq7"},
		{12812317, "H4lr"},
		{123812176, "aCVN8"},
		{123987123176, "IfKvKB2"},
		{1239871231761, "veKFNpL"},
		{62398712317612121, "ZjcNzJmmb4"},
		{623987123176121212, "gLFotDnr5k"},
		{923987123176121212, "W4PURSzrF61"},
	}

	for _, test := range tests {
		if actual := Encodei(test.value); actual != test.expected {
			t.Errorf("Encodei(%d) should give %q but received %q",
				test.value, test.expected, actual)
		}
	}
}

// Encode_fixed

func Test_Encode_fixed_7width(t *testing.T) {
	tests := []struct {
		value    int
		expected string
	}{
		{1, "1000000"},
		{76, "E100000"},
		{765, "LC00000"},
		{12837, "3L30000"},
		{987123, "Ln84000"},
		{1253667, "R8G5000"},
		{19283798, "0buI100"},
		{9817239890, "sY8OiA0"},
		{98172398912, "ucNt9j1"},
		{981723989123, "5HovaHH"},
		{2155555555555, "bBRtswb"},
		{2155555555512355, ""},
	}

	const num = 7

	for _, test := range tests {
		if actual := Encode_fixed(test.value, num); actual != test.expected {
			t.Errorf("Encode_fixed(%d, %d) should give %q but received %q",
				test.value, num, test.expected, actual)
		}
	}
}

// Benchmark Encode

func Benchmark_Encode_One(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("1")
	}
}

func Benchmark_Encode_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("1000")
	}
}

func Benchmark_Encode_953(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("953")
	}
}

func Benchmark_Encode_765234987234(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("765234987234")
	}
}

func Benchmark_Encode_12345678901234567890(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("12345678901234567890")
	}
}

// Benchmark Encodei

func Benchmark_Encodei_One(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encodei(1)
	}
}

func Benchmark_Encodei_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encodei(1000)
	}
}

func Benchmark_Encodei_953(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encodei(953)
	}
}

func Benchmark_Encodei_765234987234(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encodei(765234987234)
	}
}

func Benchmark_Encodei_1234567890123456789(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encodei(1234567890123456789)
	}
}
