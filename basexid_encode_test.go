package basex

import "testing"

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

// Encodei

func Test_Encodei_Zero(t *testing.T) {
	test("", Encodei(0), t)
}

func Test_Encodei_One(t *testing.T) {
	test("1", Encodei(1), t)
}

func Test_Encodei1(t *testing.T) {
	test("1w", Encodei(120), t)
}

func Test_Encodei2(t *testing.T) {
	test("pZN", Encodei(198237), t)
}

func Test_Encodei3(t *testing.T) {
	test("2BKvKfI", Encodei(123987123176), t)
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
