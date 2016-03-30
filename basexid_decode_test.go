package basex

import "testing"

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
