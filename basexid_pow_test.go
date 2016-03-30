package basex

import "testing"

func Test_pow1(t *testing.T) {
	test_int(8, pow(2, 3), t)
}

func Test_pow2(t *testing.T) {
	test_int(27, pow(3, 3), t)
}

func Test_pow3(t *testing.T) {
	test_int(10000, pow(10, 4), t)
}

func Test_pow4(t *testing.T) {
	test_int(1, pow(1, 0), t)
}

func Test_pow5(t *testing.T) {
	test_int(65536, pow(16, 4), t)
}
