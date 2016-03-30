package basex

import "testing"

func test(exp, act string, t *testing.T) {
	if exp != act {
		t.Errorf("Expected: '%s', got '%s'", exp, act)
	}
}

func Test_Encode_Zero(t *testing.T) {
	test("", Encode("0"), t)
}

func Test_Encode_One(t *testing.T) {
	test("1", Encode("1"), t)
}

func Test_Encode1(t *testing.T) {
	test("1w", Encode("120"), t)
}

func Test_Encode2(t *testing.T) {
	test("pZN", Encode("198237"), t)
}

func Test_Encode3(t *testing.T) {
	test("2BKvKfI", Encode("123987123176"), t)
}

// Decode

func Test_Decode_Zero(t *testing.T) {
	test("0", Decode(""), t)
}

func Test_Decode_One(t *testing.T) {
	test("1", Decode("1"), t)
}

func Test_Decode1(t *testing.T) {
	test("120", Decode("1w"), t)
}

func Test_Decode2(t *testing.T) {
	test("198237", Decode("pZN"), t)
}

func Test_Decode3(t *testing.T) {
	test("123987123176", Decode("2BKvKfI"), t)
}

// Benchmarks

func Benchmark_One(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("1")
	}
}

func Benchmark_1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("1000")
	}
}

func Benchmark_953(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("953")
	}
}

func Benchmark_765234987234(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("765234987234")
	}
}

func Benchmark_12345678901234567890(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode("12345678901234567890")
	}
}
